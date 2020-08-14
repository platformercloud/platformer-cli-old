package cli

// Error must implemented by all cli errors in this package
type Error interface {
	Error() string
	printAndExit()
}

// HandleErrorAndExit can wrap all top level command functions.
// Note: HandleErrorAndExit does not print 'Usage'
// Example: HandleErrorAndExit(logOut())
func HandleErrorAndExit(err error) {
	if err == nil {
		return
	}

	switch e := err.(type) {
	case Error:
		e.printAndExit()
		break
	default:
		// Unspecified error: will be cast as an internal error
		err := &InternalError{err, "refer debug-logs"}
		err.printAndExit()
	}
}

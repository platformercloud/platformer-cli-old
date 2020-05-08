package cli

import (
	"os"
)

// CancelError is thrown when an action is cancelled by the User.
// Does not display a message. Exits with code 0.
type CancelError struct{}

func (e *CancelError) Error() string {
	return ""
}

// HandleAndExit prints a message to the user to log in
func (e *CancelError) printAndExit() {
	os.Exit(0)
}

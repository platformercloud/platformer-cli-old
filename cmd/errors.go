package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"gitlab.platformer.com/project-x/platformer-cli/internal/config"
)

var redFprint = color.New(color.Bold, color.FgRed).FprintfFunc()
var magentaFprint = color.New(color.FgMagenta).FprintfFunc()

// UserError defines an error that can be safely printed
// to the console without leaking implementation details.
type UserError struct {
	error
}

// HandleAndExit prints the user-friendly error to the debug-log and prints
// a prefixed error details to stderr and exits with a non-zero return.
func (e *UserError) HandleAndExit() {
	log.Printf("[user] error: %v\n", e)
	redFprint(os.Stderr, "Error: %s\n", e)
	os.Exit(1)
}

// InternalError defines an error that cannot be
// directly printed to the user. Use the Message paramter
// to define a user-friendly error.
type InternalError struct {
	error
	Message string
}

// HandleAndExit prints the actual error to the debug-log and prints
// a prefixed user friendly 'Message' to stderr and exits with a non-zero return.
func (e *InternalError) HandleAndExit() {
	log.Printf("[internal] error: %v\n", e)

	redFprint(os.Stderr, "An unexpected error has occured: %s\n", e.Message)
	magentaFprint(os.Stderr, "This is likely a problem with the Platformer CLI.\n")
	fmt.Fprintf(
		os.Stderr, "Refer to the debug-log at %s and contact support through "+
			"https://platformer.com/contact\n",
		config.LogFile,
	)
	os.Exit(1)
}

// HandleErrorAndExit wraps all top level command functions.
// Example: HandleErrorAndExit(logOut())
func HandleErrorAndExit(err error) {
	if err == nil {
		return
	}

	switch e := err.(type) {
	case UserError:
		e.HandleAndExit()
	case InternalError:
		e.HandleAndExit()
	default:
		// Unspecified error: will be cast as an internal error
		err := InternalError{err, "refer debug-logs"}
		err.HandleAndExit()
	}
}

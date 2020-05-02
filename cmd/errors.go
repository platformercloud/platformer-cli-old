package cmd

import (
	"fmt"
	"log"
	"os"

	"gitlab.platformer.com/project-x/platformer-cli/internal/config"
)

// UserError defines an error that can be safely
// printed to the console without leaking implementation
// details.
type UserError struct {
	error
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
	fmt.Fprintf(
		os.Stderr,
		"An unexpected error has occured: %s\n"+
			"This is likely a problem with the Platformer CLI.\n"+
			"Refer to the debug-log at %s and contact support through:"+
			"https://platformer.com/contact\n",
		e.Message, config.LogFile,
	)
	os.Exit(1)
}

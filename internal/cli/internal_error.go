package cli

import (
	"fmt"
	"log"
	"os"

	"gitlab.platformer.com/project-x/platformer-cli/internal/config"
)

// InternalError implements a cli.Error that cannot be
// be used to log an internal error and print a user-friendly
// message to the CLI user.
type InternalError struct {
	// The actual error
	Err error

	// User-friendly message to print to the user
	// instead of the actual error
	Message string
}

func (e *InternalError) printAndExit() {
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

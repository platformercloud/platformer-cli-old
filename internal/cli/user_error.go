package cli

import (
	"fmt"
	"os"
)

// UserError implements a cli.Error that can be safely printed
// to the console without leaking implementation details.
type UserError struct{ Message string }

func (e *UserError) Error() string {
	return fmt.Sprintf("%s: %s", redSprint("Error:"), e.Message)
}

func (e *UserError) printAndExit() {
	fmt.Fprintln(os.Stderr, e)
	os.Exit(1)
}

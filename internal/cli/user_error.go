package cli

import (
	"os"

	"github.com/gookit/color"
)

// UserError implements a cli.Error that can be safely printed
// to the console without leaking implementation details.
type UserError struct{ Message string }

func (e *UserError) Error() string {
	return e.Message
}

func (e *UserError) printAndExit() {
	color.Error.Tips(e.Message)
	os.Exit(1)
}

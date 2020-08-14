package cli

import (
	"fmt"
	"os"

	"github.com/gookit/color"
)

// NotLoggedInError is thrown when the user is not logged
// into the CLI
type NotLoggedInError struct{ Message string }

func (e *NotLoggedInError) Error() string {
	return fmt.Sprintf("%s: %s", color.Danger.Render("Authentication Error"), e.Message)
}

// HandleAndExit prints a message to the user to log in
func (e *NotLoggedInError) printAndExit() {
	fmt.Fprintln(os.Stderr, e)
	fmt.Println("Use `platformer login` to log in and then retry the this command")
	os.Exit(1)
}

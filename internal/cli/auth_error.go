package cli

import (
	"fmt"
	"os"

	"github.com/gookit/color"
)

// NotLoggedInError is thrown when the user is not logged
// into the CLI
type NotLoggedInError struct{}

func (e *NotLoggedInError) Error() string {
	return ""
}

// HandleAndExit prints a message to the user to log in
func (e *NotLoggedInError) printAndExit() {
	color.Error.Tips("You are not logged in.")
	fmt.Println("Use `platformer login` to log in first and then retry the this command.")
	os.Exit(1)
}

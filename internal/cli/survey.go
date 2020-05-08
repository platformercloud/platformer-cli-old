package cli

import (
	"errors"

	"github.com/AlecAivazis/survey/v2/terminal"
)

// TransformSurveyError transforms a survey/v2 error to an cli.Error
func TransformSurveyError(err error) Error {
	if errors.As(err, &terminal.InterruptErr) {
		return &CancelError{}
	}
	// Not an interupt error; return it to the user, it's "most likely" a user error.
	return &UserError{Message: err.Error()}
}

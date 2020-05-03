package selectprompt

import (
	"github.com/spf13/cobra"
	"gitlab.platformer.com/project-x/platformer-cli/internal/cli"
)

var projectSelectCmd = &cobra.Command{
	Use:     "project",
	Aliases: []string{"project", "proj", "projs"},
	Run: func(cmd *cobra.Command, args []string) {
		cli.HandleErrorAndExit(selectProjectPrompt())
	},
}

func selectProjectPrompt() error {
	return nil
	// projectNames = auth.

	// prompt := promptui.Select{
	// 	Label: "Select Project",
	// 	Items: projectNames,
	// }

	// _, result, err := prompt.Run()
}

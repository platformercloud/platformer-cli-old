package selectprompt

import (
	"github.com/platformer-com/platformer-cli/internal/auth"
	"github.com/platformer-com/platformer-cli/internal/cli"
	"github.com/spf13/cobra"
)

// SelectCmd is the base command for all resource select prompt commands
var SelectCmd = &cobra.Command{
	Use:   "select",
	Short: "Open a select prompt to set your default organization and project",
	ValidArgs: []string{
		organizationSelectCmd.Use,
		projectSelectCmd.Use,
	},
	ArgAliases: append(
		organizationSelectCmd.Aliases,
		projectSelectCmd.Aliases...,
	),
	Args: cobra.ExactValidArgs(1),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Runs before all child commands (eg. project/org select)
		cli.HandleErrorAndExit(func() error {
			if !auth.IsLoggedIn() {
				return &cli.NotLoggedInError{}
			}
			return nil
		}())
	},
}

func init() {
	SelectCmd.AddCommand(organizationSelectCmd)
	SelectCmd.AddCommand(projectSelectCmd)
}

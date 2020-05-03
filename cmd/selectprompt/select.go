package selectprompt

import (
	"github.com/spf13/cobra"
	"gitlab.platformer.com/project-x/platformer-cli/internal/auth"
	"gitlab.platformer.com/project-x/platformer-cli/internal/cli"
)

// SelectCmd is the base command for all resource select prompt commands
var SelectCmd = &cobra.Command{
	Use:   "select",
	Short: "Open a select prompt to set your default organization and project",
	ValidArgs: []string{
		organizationSelectCmd.Use,
	},
	ArgAliases: append(
		organizationSelectCmd.Aliases,
	),
	Args: cobra.ExactValidArgs(1),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Runs before all child commands (eg. project/org list)
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
}

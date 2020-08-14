package list

import (
	"github.com/platformercloud/platformer-cli/internal/auth"
	"github.com/platformercloud/platformer-cli/internal/cli"
	"github.com/spf13/cobra"
)

// ListCmd is the base command for all resource list commands
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Platformer Resources",
	ValidArgs: []string{
		projectListCmd.Use,
		organizationListCmd.Use,
	},
	ArgAliases: append(
		projectListCmd.Aliases,
		organizationListCmd.Aliases...,
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
	// Register sub commands for the list command
	ListCmd.AddCommand(organizationListCmd)
	ListCmd.AddCommand(projectListCmd)
}

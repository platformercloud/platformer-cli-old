package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.platformer.com/project-x/platformer-cli/internal/auth"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
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
		if !auth.IsLoggedIn() {
			err := NotLoggedInError{}
			err.HandleAndExit()
		}
	},
}

func init() {
	// Register subcommands for the list command
	listCmd.AddCommand(organizationListCmd)
	listCmd.AddCommand(projectListCmd)

	rootCmd.AddCommand(listCmd)
}

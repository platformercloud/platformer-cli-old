package cmd

import (
	"fmt"

	"github.com/platformer-com/platformer-cli/internal/auth"
	"github.com/platformer-com/platformer-cli/internal/cli"
	"github.com/platformer-com/platformer-cli/internal/config"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out from your Platformer account",
	Run: func(cmd *cobra.Command, args []string) {
		cli.HandleErrorAndExit(logOut())
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}

func logOut() error {
	if !auth.IsLoggedIn() {
		return &cli.NotLoggedInError{}
	}

	config.RemoveToken()
	fmt.Println("Successfully logged out from Platformer Cloud")
	return nil
}

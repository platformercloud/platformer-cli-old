package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.platformer.com/project-x/platformer-cli/internal/auth"
	"gitlab.platformer.com/project-x/platformer-cli/internal/config"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out from your Platformer account",
	Run: func(cmd *cobra.Command, args []string) {
		HandleErrorAndExit(logOut())
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}

func logOut() error {
	if !auth.IsLoggedIn() {
		return &UserError{fmt.Errorf("you are not logged in")}
	}

	config.RemoveToken()
	fmt.Println("Successfully logged out from Platformer Cloud")
	return nil
}

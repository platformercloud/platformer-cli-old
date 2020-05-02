package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.platformer.com/project-x/platformer-cli/internal/auth"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out from your Platformer account",
	Run: func(cmd *cobra.Command, args []string) {
		HandleCommandAndExit(logOut())
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}

func logOut() error {
	return UserError{fmt.Errorf("fuck me")}
	if !auth.IsLoggedIn() {
		fmt.Println("Not logged in")
		return fmt.Errorf("not logged in")
	}

	auth.RemoveToken()
	fmt.Println("Successfully logged out from Platformer Cloud")
	return nil
}

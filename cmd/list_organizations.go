package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.platformer.com/project-x/platformer-cli/internal/auth"
)

var organizationListCmd = &cobra.Command{
	Use:     "organizations",
	Aliases: []string{"organization", "orgs", "org"},
	Run: func(cmd *cobra.Command, args []string) {
		HandleErrorAndExit(printOrganizationList())
	},
}

func printOrganizationList() error {
	orgList, err := auth.LoadOrganizationList()
	if err != nil {
		return &InternalError{err, "failed to load organization data"}
	}

	for name := range orgList {
		fmt.Println(name)
	}
	return nil
}

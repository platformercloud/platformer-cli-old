package list

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.platformer.com/project-x/platformer-cli/internal/auth"
	"gitlab.platformer.com/project-x/platformer-cli/internal/cli"
)

var (
	organizationListCmd = &cobra.Command{
		Use:     "organizations",
		Aliases: []string{"organization", "org", "orgs"},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := preloadOrganizationsAndValidateSelected(); err != nil {
				// Returning an error to *RunE() will print the 'Usage' as well.
				return err
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			cli.HandleErrorAndExit(printOrgList())
		},
	}
)

func printOrgList() error {
	orgList, err := auth.LoadOrganizationList()
	if err != nil {
		return &cli.InternalError{Err: err, Message: "failed to load organization data"}
	}

	for orgName := range orgList {
		fmt.Println(orgName)
	}

	return nil
}

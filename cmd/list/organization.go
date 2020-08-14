package list

import (
	"fmt"
	"github.com/platformercloud/platformer-cli/internal/auth"
	"github.com/platformercloud/platformer-cli/internal/cli"
	"strings"

	"github.com/spf13/cobra"
)

var organizationListCmd = &cobra.Command{
	Use:     "organizations",
	Aliases: []string{"organization", "org", "orgs"},
	Run: func(cmd *cobra.Command, args []string) {
		cli.HandleErrorAndExit(printOrgList())
	},
}

func printOrgList() error {
	orgList, err := auth.LoadOrganizationList()
	if err != nil {
		return &cli.InternalError{Err: err, Message: "failed to load organization data"}
	}

	fmt.Println(strings.Join(orgList.Names(), "\n"))
	return nil
}

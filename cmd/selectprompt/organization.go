package selectprompt

import (
	"github.com/spf13/cobra"
	"gitlab.platformer.com/project-x/platformer-cli/internal/cli"
)

var organizationSelectCmd = &cobra.Command{
	Use:     "organization",
	Aliases: []string{"organizations", "orgs", "org"},
	Run: func(cmd *cobra.Command, args []string) {
		cli.HandleErrorAndExit(printOrganizationList())
	},
}

package selectprompt

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"gitlab.platformer.com/project-x/platformer-cli/internal/auth"
	"gitlab.platformer.com/project-x/platformer-cli/internal/cli"
	"gitlab.platformer.com/project-x/platformer-cli/internal/config"
)

var organizationSelectCmd = &cobra.Command{
	Use:     "organization",
	Aliases: []string{"organizations", "orgs", "org"},
	Run: func(cmd *cobra.Command, args []string) {
		cli.HandleErrorAndExit(selectOrgWithPrompt())
	},
}

func selectOrgWithPrompt() error {
	orgList, err := auth.LoadOrganizationList()
	if err != nil {
		return &cli.InternalError{Err: err, Message: "failed to load organization data"}
	}

	prompt := promptui.Select{
		Label: "Select Organization",
		Items: orgList.Names(),
	}

	_, selectedOrgName, err := prompt.Run()
	if err != nil {
		// Do nothing; this error is thrown if the user quits the CLI with ctrl+C
	}

	if selectedOrgName == "" {
		return &cli.UserError{Message: "no organization has been selected"}
	}

	config.SetDefaultOrg(selectedOrgName)

	green := color.New(color.FgHiGreen).SprintfFunc()
	fmt.Printf("%s has been set as your default organization\n", green(selectedOrgName))

	return nil
}

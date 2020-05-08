package selectprompt

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
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

	prompt := &survey.Select{
		Message: "Select Organization",
		Options: orgList.Names(),
	}

	var selectedOrgName string
	if err := survey.AskOne(prompt, &selectedOrgName); err != nil {
		return cli.TransformSurveyError(err)
	}

	if selectedOrgName == "" {
		return &cli.UserError{Message: "no organization has been selected"}
	}

	config.SetDefaultOrg(selectedOrgName)

	green := color.New(color.FgHiGreen).SprintfFunc()
	fmt.Printf("%s has been set as your default organization\n", green(selectedOrgName))

	return nil
}

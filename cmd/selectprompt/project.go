package selectprompt

import (
	"fmt"
	"github.com/platformercloud/platformer-cli/internal/auth"
	"github.com/platformercloud/platformer-cli/internal/cli"
	"github.com/platformercloud/platformer-cli/internal/config"

	"github.com/AlecAivazis/survey/v2"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var (
	// organization name flag
	// overrides the default/saved organization if provided
	orgNameFlag string

	projectSelectCmd = &cobra.Command{
		Use:     "project",
		Aliases: []string{"project", "proj", "projs"},
		Run: func(cmd *cobra.Command, args []string) {
			cli.HandleErrorAndExit(selectProjectWithPrompt())
		},
	}
)

func init() {
	projectSelectCmd.Flags().StringVarP(&orgNameFlag, "organization", "o", "", "--organization=<ORG_NAME> or -o <ORG_NAME>")
}

func selectProjectWithPrompt() error {
	orgName, projectList, err := auth.LoadProjectsFromDefaultOrFlag(orgNameFlag)
	if err != nil {
		return err
	}

	prompt := &survey.Select{
		Message: "Select Organization",
		Options: projectList.Names(),
	}

	var selectedProjectName string
	if err := survey.AskOne(prompt, &selectedProjectName); err != nil {
		return cli.TransformSurveyError(err)
	}

	if selectedProjectName == "" {
		return &cli.UserError{Message: "no organization has been selected"}
	}

	config.SetDefaultProject(selectedProjectName)

	green := color.FgLightGreen.Render
	fmt.Printf("%s and %s has been set as your default Organization and Project\n",
		green(orgName),
		green(selectedProjectName),
	)

	return nil
}

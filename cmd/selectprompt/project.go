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

	prompt := promptui.Select{
		Label: "Select Project from " + orgName,
		Items: projectList.Names(),
	}

	_, selectedProjectName, err := prompt.Run()
	if err != nil {
		// Do nothing; this error is thrown if the user quits the CLI with ctrl+C
	}

	if selectedProjectName == "" {
		return &cli.UserError{Message: "no project has been selected"}
	}

	config.SetDefaultProject(selectedProjectName)

	green := color.New(color.FgHiGreen).SprintfFunc()
	fmt.Printf("%s > %s has been set as your default organization and project\n",
		green(orgName),
		green(selectedProjectName),
	)

	return nil
}

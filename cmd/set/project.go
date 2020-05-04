package set

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"gitlab.platformer.com/project-x/platformer-cli/internal/auth"
	"gitlab.platformer.com/project-x/platformer-cli/internal/cli"
	"gitlab.platformer.com/project-x/platformer-cli/internal/config"
)

var (
	// organization name flag
	// overrides the default/saved organization if provided
	orgNameFlag string

	projectSetCmd = &cobra.Command{
		Use:     "project",
		Aliases: []string{"proj"},
		Args:    cobra.ExactArgs(1), // requires exactly 1 argument (project name)
		Run: func(cmd *cobra.Command, args []string) {
			projectName := args[0]
			cli.HandleErrorAndExit(validateAndSetProject(projectName))
		},
	}
)

func init() {
	projectSetCmd.Flags().StringVarP(&orgNameFlag, "organization", "o", "", "--organization=<ORG_NAME> or -o <ORG_NAME>")
}

func validateAndSetProject(projectName string) error {
	orgName, projectList, err := auth.LoadProjectsFromDefaultOrFlag(orgNameFlag)
	if err != nil {
		return err
	}

	if _, ok := projectList[projectName]; !ok {
		return &cli.UserError{"invalid Project name: " + projectName}
	}

	config.SetDefaultProject(projectName)

	green := color.New(color.FgHiGreen).SprintfFunc()
	fmt.Printf("%s > %s has been set as your default organization and project\n",
		green(orgName),
		green(projectName),
	)

	return nil
}

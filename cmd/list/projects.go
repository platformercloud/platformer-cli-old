package list

import (
	"fmt"
	"github.com/platformercloud/platformer-cli/internal/auth"
	"github.com/platformercloud/platformer-cli/internal/cli"
	"strings"

	"github.com/spf13/cobra"
)

var (
	// organization name flag
	// overrides the default/saved organization if provided
	orgNameFlag string

	projectListCmd = &cobra.Command{
		Use:     "projects",
		Aliases: []string{"project", "proj", "projs"},
		Run: func(cmd *cobra.Command, args []string) {
			cli.HandleErrorAndExit(printProjectList())
		},
	}
)

func init() {
	projectListCmd.Flags().StringVarP(&orgNameFlag, "organization", "o", "", "--organization=<ORG_NAME> or -o <ORG_NAME>")
}

func printProjectList() error {
	_, projectList, err := auth.LoadProjectsFromDefaultOrFlag(orgNameFlag)
	if err != nil {
		return err
	}

	fmt.Println(strings.Join(projectList.Names(), "\n"))
	return nil
}

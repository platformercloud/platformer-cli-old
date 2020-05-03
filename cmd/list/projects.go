package list

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"gitlab.platformer.com/project-x/platformer-cli/internal/auth"
	"gitlab.platformer.com/project-x/platformer-cli/internal/cli"
	"gitlab.platformer.com/project-x/platformer-cli/internal/config"
)

var (
	// organization name flag
	orgNameFlag string

	orgList     auth.OrganizationList
	selectedOrg auth.Organization

	projectListCmd = &cobra.Command{
		Use:     "projects",
		Aliases: []string{"project", "proj", "projs"},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := preloadOrganizationsAndValidateSelected(); err != nil {
				// Returning an error to *RunE() will print the 'Usage' as well.
				return err
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			cli.HandleErrorAndExit(printProjectList())
		},
	}
)

func init() {
	projectListCmd.Flags().StringVarP(&orgNameFlag, "organization", "o", "", "--organization=<ORG_NAME> or -o <ORG_NAME>")
}

func preloadOrganizationsAndValidateSelected() (err error) {
	orgList, err = auth.LoadOrganizationList()
	if err != nil {
		return &cli.InternalError{Err: err, Message: "failed to load organization data"}
	}

	var orgNameKey string
	if orgNameFlag != "" {
		// The org flag has been set; ignore the currently selected org in the local config
		if _, ok := orgList[orgNameFlag]; !ok {
			return &cli.UserError{Message: fmt.Sprintf("the organization [%s] does not exist "+
				"or you do not have permission to access it", orgNameFlag)}
		}
		orgNameKey = orgNameFlag
	} else if defaultOrgName := config.GetDefaultOrg(); defaultOrgName != "" {
		if _, ok := orgList[defaultOrgName]; !ok {
			return &cli.UserError{Message: fmt.Sprintf("the currently selected organization [%s] does not exist "+
				"or you do not have permission to access it."+
				"\nUse --organization=<ORG_NAME> or `platformer use organization` to set a valid organization",
				orgNameFlag),
			}
		}
		orgNameKey = defaultOrgName
	}

	if orgNameKey == "" {
		return &cli.UserError{Message: "an organization must be selected." +
			"\nUse --organization=<ORG_NAME> or `platformer use organization` to set a valid organization",
		}
	}

	selectedOrg = orgList[orgNameKey]
	return nil
}

func printProjectList() error {
	projectList, err := auth.LoadProjectList(selectedOrg.ID)
	if err != nil {
		return &cli.InternalError{Err: err, Message: "failed to load project data"}
	}

	fmt.Println(strings.Join(projectList.Names(), "\n"))
	return nil
}

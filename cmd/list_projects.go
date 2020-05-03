package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.platformer.com/project-x/platformer-cli/internal/auth"
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
			HandleErrorAndExit(printProjectList())
		},
	}
)

func init() {
	projectListCmd.Flags().StringVarP(&orgNameFlag, "organization", "o", "", "--organization=<ORG_NAME> or -o <ORG_NAME>")
	projectListCmd.MarkFlagRequired("organization")
}

func preloadOrganizationsAndValidateSelected() (err error) {
	orgList, err = auth.LoadOrganizationList()
	if err != nil {
		return &InternalError{err, "failed to load organization data"}
	}

	var orgNameKey string
	if orgNameFlag != "" {
		// The org flag has been set; ignore the currently selected org in the local config
		if _, ok := orgList[orgNameFlag]; !ok {
			return &UserError{
				fmt.Errorf("the organization [%s] does not exist or you do not have permission to access it", orgNameFlag),
			}
		}
		orgNameKey = orgNameFlag
	} else if savedOrgName := viper.GetString("context.organization.name"); savedOrgName != "" {
		if _, ok := orgList[savedOrgName]; !ok {
			return &UserError{
				fmt.Errorf("the currently selected organization [%s] does not exist"+
					"or you do not have permission to access it."+
					"\nUse --organization=<ORG_NAME> or `platformer use organization` to set a valid organization",
					orgNameFlag),
			}
		}
		orgNameKey = savedOrgName
	}

	if orgNameKey == "" {
		return &UserError{
			fmt.Errorf("an organization must be selected." +
				"\nUse --organization=<ORG_NAME> or `platformer use organization` to set a valid organization",
			),
		}
	}

	selectedOrg = orgList[orgNameKey]
	return nil
}

func printProjectList() error {
	projectList, err := auth.LoadProjectList(selectedOrg.ID)
	if err != nil {
		return &InternalError{err, "failed to load project data"}
	}

	for p := range projectList {
		fmt.Println(p)
	}
	return nil
}

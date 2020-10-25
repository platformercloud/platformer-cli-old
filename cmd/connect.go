package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/gookit/color"
	"github.com/platformercloud/platformer-cli/internal/auth"
	"github.com/platformercloud/platformer-cli/internal/cli"
	"github.com/platformercloud/platformer-cli/internal/config"
	"github.com/platformercloud/platformer-cli/internal/mizzen"
	"github.com/spf13/cobra"
)

var (
	orgNameFlag     string
	projectNameFlag string

	connectCmd = &cobra.Command{
		Use:   "connect",
		Short: "Connect your Kubernetes Cluster to Platformer Cloud",
		Run: func(cmd *cobra.Command, args []string) {
			cli.HandleErrorAndExit(connectCluster())
		},
	}
)

// @todo - needs to be changed in production. These values must be injected at compile time somehow...
const consoleURL = "https://beta.console.platformer.com"

func init() {
	connectCmd.Flags().StringVarP(&orgNameFlag, "organization", "o", "", "--organization=<ORG_NAME> or -o <ORG_NAME>")
	connectCmd.Flags().StringVarP(&projectNameFlag, "project", "p", "", "--project=<PROJECT_NAME> or -p <PROJECT_NAME>")
	rootCmd.AddCommand(connectCmd)
}

func connectCluster() error {
	if orgNameFlag == "" {
		orgNameFlag = config.GetDefaultOrg()
	}

	if projectNameFlag == "" {
		projectNameFlag = config.GetDefaultProject()
	}

	organization, err := auth.GetOrganizationIDFromName(orgNameFlag)
	if err != nil {
		return &cli.UserError{Message: "invalid organization"}
	}

	project, err := auth.GetProjectIDFromName(organization.ID, projectNameFlag)
	if err != nil {
		return &cli.UserError{Message: "invalid project"}
	}

	kw, err := mizzen.NewKubectlWrapper()
	if err != nil {
		return &cli.UserError{Message: "kubectl binary not found. Please install and configure kubectl first."}
	}

	clusterList, err := getClusterList(kw)
	if err != nil {
		return err
	}

	var qs = []*survey.Question{
		{
			Name:     "clusterName",
			Validate: survey.Required,
			Prompt: &survey.Select{
				Message: "Select Cluster to connect",
				Options: clusterList,
			},
		},
	}

	var selection struct{ ClusterName string }
	if err := survey.Ask(qs, &selection); err != nil {
		return cli.TransformSurveyError(err)
	}

	green := color.FgGreen.Render
	confirmPrompt := &survey.Confirm{
		Default: true,
		Message: fmt.Sprintf("Cluster %s will be connected under Organization %s and Project %s. Proceed?",
			green(selection.ClusterName), green(orgNameFlag), green(projectNameFlag),
		),
	}

	var confirm bool
	if err := survey.AskOne(confirmPrompt, &confirm, survey.WithValidator(survey.Required)); err != nil {
		return cli.TransformSurveyError(err)
	}

	if !confirm {
		return &cli.CancelError{}
	}

	fmt.Println("Connecting your cluster... This may take a few minutes. Please do not exit the terminal.")
	if err := mizzen.ConnectAndInstallAgent(kw, organization.ID, project.ID, selection.ClusterName); err != nil {
		return err
	}

	fmt.Println(green("Success!"), "Cluster", selection.ClusterName, "has been successfully connected to Platformer Cloud")
	fmt.Println("You may now deploy applications to your application through the platformer console:", consoleURL)
	return nil
}

func getClusterList(kw *mizzen.KubectlWrapper) ([]string, cli.Error) {
	clusters, err := kw.ListClusters()
	if err != nil {
		return nil, &cli.InternalError{Message: "failed to get-clusters from kubectl", Err: err}
	}

	if len(clusters) == 0 {
		return nil, &cli.UserError{Message: "no clusters were found from kubectl"}
	}

	return clusters, nil
}

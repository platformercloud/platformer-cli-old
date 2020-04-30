/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"gitlab.platformer.com/chamod.p/platformer/internal"
	"log"
	"strings"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Set an active Platformer for your working directory",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !(len(args) > 1) {
			fmt.Println(color.YellowString("Parameters required!"))
			fmt.Println("platformer use [project/org] [name]")
			return
		}

		switch args[0] {
		case "org":
			organizationName := args[1]
			organizationID, err := getOrganizationID(organizationName)
			if err != nil {
				log.Fatalf("unable to get organization ID %s", err)
			}

			result, err := internal.SaveOrganizationGlobally(organizationID)
			if err != nil {
				fmt.Print(color.RedString("unable to save the organization"))
				fmt.Printf("%s", err)
			}

			if result {
				fmt.Println(color.GreenString("Organization Saved!"))
			}
		case "project":
			projectName := args[1]
			projectID, err := getProjectID(projectName)
			if err != nil {
				log.Fatalf("unable to get organization ID %s", err)
			}

			result, err := internal.SaveProjectGlobally(projectID)
			if err != nil {
				fmt.Print(color.RedString("Unable to save the project"))
				fmt.Printf("%s", err)
			}

			if result {
				fmt.Println(color.GreenString("Project Saved!"))
			}
		}
	},
}

func getProjectID(projectName string) (string, error) {
	projectsList, err := internal.GetProjects()
	if err != nil {
		return "", fmt.Errorf("unable to get project list %s", err)
	}

	for _, project := range projectsList {
		if strings.EqualFold(projectName, project.ProjectName) {
			return project.ProjectID, nil
		}
	}

	return "", errors.New("given project name invalid")
}

func getOrganizationID(organizationName string) (string, error) {
	organizationsList, err := internal.GetOrganizationList()
	if err != nil {
		return "", fmt.Errorf("unable to get organization list %s", err)
	}

	for _, organization := range organizationsList {
		if strings.EqualFold(organizationName, organization.Name) {
			return organization.OrganizationId, nil
		}
	}

	return "", errors.New("given organization name invalid")
}

func init() {
	rootCmd.AddCommand(useCmd)
}

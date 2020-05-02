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
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"gitlab.platformer.com/project-x/platformer-cli/internal"
)

func displayOrganizationNames(list []internal.Organization) {
	for _, organization := range list {
		fmt.Println(organization.Name)
	}
}

func displayProjectNames(list []internal.Project) {
	for _, project := range list {
		fmt.Println(project.ProjectName)
	}
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all projects or organizations ",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(color.YellowString("Parameters Required! "))
			return
		}

		switch args[0] {
		case "orgs":
			{
				orgList, err := internal.GetOrganizationList()
				if err != nil {
					log.Fatalf("unable to get organizations list %s", err)
				}
				displayOrganizationNames(orgList)
			}
		case "projects":
			projectList, err := internal.GetProjects()
			if err != nil {
				log.Fatalf("unable to get projects list %s", err)
			}
			displayProjectNames(projectList)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

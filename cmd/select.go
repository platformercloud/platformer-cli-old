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
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"gitlab.platformer.com/project-x/platformer-cli/internal"
)

// selectCmd represents the select command
var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "Select organization or project",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !(len(args) > 0) {
			fmt.Println(color.YellowString("Parameters required!"))
			fmt.Println("platformer use [project/org] [name]")
			return
		}

		switch args[0] {
		case "org":
			organizationList, err := internal.GetOrganizationList()
			if err != nil {
				log.Fatalf("%s unable to get organization list", err)
			}

			organizationNames := internal.GetOrganizationsNames(organizationList)

			prompt := promptui.Select{
				Label: "Select Organization",
				Items: organizationNames,
			}

			_, result, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			fmt.Printf("You choose %q\n", result)

			selectedOrganizationIndex := indexOf(result, organizationNames)

			isLocalConfigExists, err := internal.IsLocalConfigExists()
			if err != nil {
				log.Fatalf("error getting local config existence %s", err)
			}

			if isLocalConfigExists {
				isSaved, err := internal.SaveOrganizationLocally(organizationList[selectedOrganizationIndex].OrganizationId)
				if err != nil {
					log.Fatalf("unable to save organization locally %s", err)
				}

				if isSaved {
					fmt.Println(color.GreenString("Organization saved locally"))

					return
				}

				fmt.Println(color.RedString("Unable save the organization"))
			}

			// If Local configuration doesn't exists
			isSaved, err := internal.SaveOrganizationGlobally(organizationList[selectedOrganizationIndex].OrganizationId)
			if err != nil {
				log.Fatalf("unable to save organization globally %s", err)
			}

			if isSaved {
				fmt.Println(color.GreenString("Organization saved globally"))
				return
			}

			fmt.Println(color.RedString("Unable save the organization"))

		case "project":
			projectList, err := internal.GetProjects()
			if err != nil {
				log.Fatalf("%s unable to get project list", err)
			}

			projectNames := internal.GetProjectsNames(projectList)

			prompt := promptui.Select{
				Label: "Select Project",
				Items: projectNames,
			}

			_, result, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			fmt.Printf("You choose %q\n", result)

			selectedProjectIndex := indexOf(result, projectNames)

			isLocalConfigExists, err := internal.IsLocalConfigExists()
			if err != nil {
				log.Fatalf("error getting local config existence %s", err)
			}

			if isLocalConfigExists {
				isSaved, err := internal.SaveProjectLocally(projectList[selectedProjectIndex].ProjectID)
				if err != nil {
					log.Fatalf("unable to save project locally %s", err)
				}

				if isSaved {
					fmt.Println(color.GreenString("Project saved locally"))
					return
				}

				fmt.Println(color.RedString("Unable save the project"))
			}

			// If Local configuration doesn't exists
			isSaved, err := internal.SaveProjectGlobally(projectList[selectedProjectIndex].ProjectID)
			if err != nil {
				log.Fatalf("unable to save project globally %s", err)
			}

			if isSaved {
				fmt.Println(color.GreenString("Project saved globally"))
				return
			}

			fmt.Println(color.RedString("Unable save the project"))
		}
	},
}

func init() {
	rootCmd.AddCommand(selectCmd)
}

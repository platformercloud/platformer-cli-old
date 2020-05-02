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
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"gitlab.platformer.com/project-x/platformer-cli/internal"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		banner := `
######   #           #     #######  #######  #######  ######   #     #  #######  ######
#     #  #         #   #      #     #        #     #  #     #  # # # #  #        #     # 
######   #        #     #     #     #####    #     #  ######   #  #  #  #####    ######  
#        #        #######     #     #        #     #  #   #    #     #  #        #   #
#        #######  #     #     #     #        #######  #     #  #     #  #######  #     #`
		fmt.Println(color.BlueString(banner))
		fmt.Println("\nYou're about to initialize a Platformer project in this directory: ")

		projectList, err := internal.GetProjects()
		if err != nil {
			log.Fatalf("%s", err)
		}

		projectNamesArray := ProjectListToArray(projectList)

		prompt := promptui.Select{
			Label: "Select Organization",
			Items: projectNamesArray,
		}

		_, result, err := prompt.Run()

		if err != nil {
			log.Fatalf("Prompt failed %v\n", err)
			return
		}

		fmt.Printf("You choose %q\n", result)

		selectedProjectIndex := indexOf(result, projectNamesArray)
		err = saveProjectSetting(projectList[selectedProjectIndex])
		if err != nil {
			log.Fatalf("unable to save project settings %s", err)
		}

	},
}

func saveProjectSetting(project internal.Project) error {
	var dir string
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("unable to get working directory %s", err)
	}

	_, err = os.Stat(dir + "/.platformer/projects")
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dir+"/.platformer/", 0755)
		if errDir != nil {
			return fmt.Errorf("unable to create dir %s", errDir)
		}
	}

	f, err := os.Create(dir + "/.platformer/projects")
	if err != nil {
		return fmt.Errorf("error file creating. %s", err)
	}
	writer := bufio.NewWriter(f)
	_, err = writer.WriteString(project.ProjectID)
	if err != nil {
		return fmt.Errorf("error token wrting %s", err)
	}
	_ = writer.Flush()

	return nil
}

func indexOf(word string, data []string) int {
	for k, v := range data {
		if word == v {
			return k
		}
	}
	return -1
}

func ProjectListToArray(projectList []internal.Project) []string {

	var projectNamesList []string

	for _, project := range projectList {
		projectNamesList = append(projectNamesList, project.ProjectName)
	}

	return projectNamesList
}

func init() {
	rootCmd.AddCommand(initCmd)
}

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
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Kompose is a conversion tool for Docker Compose to container orchestrators such as Kubernetes (or OpenShift).",
	Long: `Kompose is a conversion tool for Docker Compose to container orchestrators such as Kubernetes (or OpenShift).
	`,
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "kompose":
			fmt.Println("Looking for docker-compose.yml....")
			cmd := exec.Command("kompose", "convert")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				log.Fatalf("cmd.Run() failed with %s\n", err)
			}
		}

	},
	Hidden: true,
}

func init() {
	rootCmd.AddCommand(convertCmd)
}

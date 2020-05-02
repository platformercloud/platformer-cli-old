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
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"gitlab.platformer.com/project-x/platformer-cli/internal"
)

func removeStoredToken(dir string) {

	// Validate created token file
	if _, err := os.Stat(dir + "/.platformer/token"); os.IsNotExist(err) {
		// TOKEN file does not exist
		fmt.Println(color.YellowString("You haven't logged in before"))
	}

	dat, err := ioutil.ReadFile(dir + "/.platformer/token")
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	token := string(dat)

	if len(token) == 0 {
		fmt.Println(color.YellowString("You have already logged out"))
		return
	}
	_, _ = os.Create(dir + "/.platformer/token")
	fmt.Println(color.GreenString("Logged out from Platformer Account"))

}

func logOut() {
	dir, err := internal.GetOSRootDir()
	if err != nil {
		log.Fatalf("error getting root dir : %s", err)
	}
	removeStoredToken(dir)
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "log the CLI out of Platformer Cloud",

	Run: func(cmd *cobra.Command, args []string) {
		logOut()
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}

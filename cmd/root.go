package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.platformer.com/project-x/platformer-cli/cmd/list"
	"gitlab.platformer.com/project-x/platformer-cli/cmd/selectprompt"
	"gitlab.platformer.com/project-x/platformer-cli/cmd/set"
	"gitlab.platformer.com/project-x/platformer-cli/internal/config"
)

// RootCmd represents the base command when called without any sub commands
var rootCmd = &cobra.Command{
	Use:   "platformer",
	Short: "A Fully-Manged Developer Platform with CI/CD and DevOps Tools baked in.",
	Long: `Utilize the power of Kubernetes without the hassle of maintaining your own Kubernetes infrastructure.
Create your own Cluster or deploy your applications on a shared Cluster fully-managed by Platformer - and pay for only what you use!`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	_ = rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(func() {
		viper.SetConfigFile(config.ConfigPath)
		viper.SetConfigType("yaml")

		if err := viper.ReadInConfig(); err != nil {
			fmt.Fprintf(os.Stderr, "fatal error: %v\n", err)
			os.Exit(1)
		}
	})

	registerModuleCommands()
}

// registerModuleCommands registers all packaged commands
// (these cannot register themselves with init())
func registerModuleCommands() {
	rootCmd.AddCommand(list.ListCmd)
	rootCmd.AddCommand(selectprompt.SelectCmd)
	rootCmd.AddCommand(set.SetCmd)
}

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitlab.platformer.com/project-x/platformer-cli/internal/config"

	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any sub commands
var rootCmd = &cobra.Command{
	Use:   "platformer",
	Short: "A Fully-Manged Developer Platform with CI/CD and DevOps Tools baked in.",
	Long: `Utilize the power of Kubernetes without the hassle of maintaining your own Kubernetes infrastructure.
Create your own Cluster or deploy your applications on a shared Cluster fully-managed by Platformer - and pay for only what you use!`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
		os.Exit(1)
	}
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

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.platformer.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

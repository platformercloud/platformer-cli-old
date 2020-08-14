package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
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

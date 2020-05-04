package main

import (
	"fmt"
	"os"

	"gitlab.platformer.com/project-x/platformer-cli/cmd"
	"gitlab.platformer.com/project-x/platformer-cli/internal/config"
)

func main() {
	if err := config.InitPlatformerDirectory(); err != nil {
		fmt.Fprintf(os.Stderr, "Fatal Error: %v\n", err)
		os.Exit(1)
	}

	if err := config.InitDebugLogs(); err != nil {
		fmt.Fprintf(os.Stderr, "Fatal Error: %v\n", err)
		os.Exit(1)
	}

	cmd.Execute()
}

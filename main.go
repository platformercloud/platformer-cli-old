package main

import (
	"fmt"
	"github.com/platformercloud/platformer-cli/cmd"
	"github.com/platformercloud/platformer-cli/internal/config"
	"io"
	"log"
	"os"
)

func main() {
	if err := config.InitPlatformerDirectory(); err != nil {
		fmt.Fprintf(os.Stderr, "Fatal Error: %v\n", err)
		os.Exit(1)
	}

	f, err := config.InitDebugLogFile()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal Error: %v\n", err)
		os.Exit(1)
	}

	defer f.Close()
	log.SetOutput(io.Writer(f))

	cmd.Execute()
}

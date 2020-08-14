package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/platformer-com/platformer-cli/cmd"
	"github.com/platformer-com/platformer-cli/internal/config"
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

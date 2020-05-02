package config

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
)

var (
	// PlatformerDir is the filepath to the .platformer directory
	PlatformerDir string

	// ConfigPath is the filepath to the Configuration (yaml) file
	ConfigPath string

	// LogFile is the filepath to debug-logs
	LogFile string
)

func init() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal: cannot find home directory: %v", err)
		os.Exit(1)
	}

	PlatformerDir = path.Join(home, ".platformer")
	ConfigPath = path.Join(PlatformerDir, "config")
	LogFile = path.Join(PlatformerDir, "debug-log")
}

// InitPlatformerDirectory prepares the .platformer directory
// structure required by the CLI. Creates new folders and files if
// they do not exist.
func InitPlatformerDirectory() error {
	var err error
	if _, err = os.Stat(ConfigPath); err == nil {
		return nil
	}

	if !os.IsNotExist(err) {
		return fmt.Errorf("unexpected error: %w", err)
	}

	// Config file not found; initialize new directory
	if err := os.Mkdir(PlatformerDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create %s directory: %w", PlatformerDir, err)
	}

	// Create new config file
	if _, err = os.Create(ConfigPath); err != nil {
		return fmt.Errorf("failed to create new config file %s: %w", ConfigPath, err)
	}

	return nil
}

// InitDebugLogs initializes the debug-log file. Appends if exists
// and sets the stdlib log output to the LogFile
func InitDebugLogs() error {
	f, err := os.OpenFile(LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("error opening debug-log file: %w", err)
	}
	defer f.Close()
	log.SetOutput(io.Writer(f))

	return nil
}

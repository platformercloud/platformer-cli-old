package internal

import (
	"errors"
	"fmt"
	"os"
	"runtime"
)

func GetOSRootDir() (string, error) {
	var dir string

	switch osType := runtime.GOOS; osType {
	case "windows":
		dir = os.Getenv("LocalAppData")
		if dir == "" {
			err := errors.New("%AppData% is not defined")
			return "", fmt.Errorf("%w", err)
		}
	case "darwin":
		dir = os.Getenv("HOME")
		if dir == "" {
			err := errors.New("$home is not defined")
			return "", fmt.Errorf("%w", err)
		}
	case "linux":
		fmt.Println("Linux Geek")
		dir = os.Getenv("home")
		if dir == "" {
			err := errors.New("$home is not defined")
			return "", fmt.Errorf("%w", err)
		}
	default:
		dir = os.Getenv("XDG_CONFIG_HOME")
		if dir == "" {
			err := errors.New("$home is not defined")
			return "", fmt.Errorf("%w", err)
		}
	}

	return dir, nil
}

package cli

import (
	"os"

	"golang.org/x/sys/windows"
)

// EnableWindowsColorSupport is cross-platform compatible.
// Enables ANSI colors on windows terminals.
func EnableWindowsColorSupport() {
	stdout := windows.Handle(os.Stdout.Fd())
	var originalMode uint32
	windows.GetConsoleMode(stdout, &originalMode)
	windows.SetConsoleMode(stdout, originalMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
}

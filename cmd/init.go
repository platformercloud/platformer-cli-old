package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"golang.org/x/sys/windows"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		banner := `
######   #           #     #######  #######  #######  ######   #     #  #######  ######
#     #  #         #   #      #     #        #     #  #     #  # # # #  #        #     # 
######   #        #     #     #     #####    #     #  ######   #  #  #  #####    ######  
#        #        #######     #     #        #     #  #   #    #     #  #        #   #
#        #######  #     #     #     #        #######  #     #  #     #  #######  #     #`
		fmt.Println(color.BlueString(banner))
		fmt.Println("\nYou're about to initialize a Platformer project in this directory: ")
	},
}

func init() {
	enableWindowsColorSupport()
	rootCmd.AddCommand(initCmd)
}

// Cross-platformer compatible, enables ANSI colors on windows terminals.
func enableWindowsColorSupport() {
	stdout := windows.Handle(os.Stdout.Fd())
	var originalMode uint32
	windows.GetConsoleMode(stdout, &originalMode)
	windows.SetConsoleMode(stdout, originalMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
}

package set

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/platformercloud/platformer-cli/internal/cli"
	"github.com/platformercloud/platformer-cli/internal/config"
	"github.com/spf13/cobra"
)

var (
	contextSetCmd = &cobra.Command{
		Use:     "context",
		Aliases: []string{"ctx"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			contextUrl := args[0]
			cli.HandleErrorAndExit(validateAndSetContextURL(contextUrl))
		},
	}
)

func validateAndSetContextURL(contextUrl string) error {

	if contextUrl != "" {
		config.SetDefaultContextURL(contextUrl)
	}

	green := color.FgLightGreen.Render
	fmt.Printf("%s has been set as your default context URL\n", green(contextUrl))

	return nil
}
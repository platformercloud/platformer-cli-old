package set

import (
	"github.com/platformercloud/platformer-cli/internal/auth"
	"github.com/platformercloud/platformer-cli/internal/cli"
	"github.com/spf13/cobra"
)

func getSetArgAliases() []string {
	aliases := [][]string {
		organizationSetCmd.Aliases,
		projectSetCmd.Aliases,
		contextSetCmd.Aliases,
	}

	concat := make([]string, len(aliases))
	for _, a  := range aliases{
		concat = append(concat, a...)
	}

	return concat
}

// SetCmd is the base command for all resource set commands
// set is the same as 'select prompt' but without the prompt list.
var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set your default organization and project",
	ValidArgs: []string{
		organizationSetCmd.Use,
		projectSetCmd.Use,
		contextSetCmd.Use,
	},
	ArgAliases: getSetArgAliases(),
	Args: cobra.ExactValidArgs(1),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Runs before all child commands (eg. project/org list)
		cli.HandleErrorAndExit(func() error {
			if !auth.IsLoggedIn() {
				return &cli.NotLoggedInError{}
			}
			return nil
		}())
	},
}

func init() {
	SetCmd.AddCommand(organizationSetCmd)
	SetCmd.AddCommand(projectSetCmd)
	SetCmd.AddCommand(contextSetCmd)
}

package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (c *CLI) newRootCommand() *cobra.Command {
	root := &cobra.Command{
		Use:   "gamesync",
		Short: "GameSync synchronizes game mods and configs",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("GameSync")
			fmt.Println(
				"Use 'gamesync --help' to see available commands.",
			)
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	root.AddCommand(
		c.newUpdateManifestCommand(),
	)

	root.AddCommand(
		c.newServeCommand(),
	)

	root.AddCommand(
		c.newSyncCommand(),
	)

	return root
}

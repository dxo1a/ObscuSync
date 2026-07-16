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
	}

	root.AddCommand(
		c.newScanCommand(),
	)

	root.AddCommand(
		c.newServeCommand(),
	)

	return root
}

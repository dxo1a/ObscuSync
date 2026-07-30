package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	cobra.MousetrapHelpText = ""
}

func (c *CLI) newRootCommand() *cobra.Command {
	root := &cobra.Command{
		Use:   "ObscuSyncCLI",
		Short: "[#] ObscuSync is a tool for syncing game mods and configurations.\n[#]   If you haven't done this yet, edit config.yaml. It was created near ObscuSyncCLI.exe file.\n[#]   Supported games are described in the same config. Maybe in the future I will remove the boundaries of supported games, but not now.\n[c] axvore/dxo1a",
		CompletionOptions: cobra.CompletionOptions{
			HiddenDefaultCmd: true,
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("ObscuSync is a tool for syncing game mods and configurations.")
			fmt.Println(
				"Use 'ObscuSyncCLI --help' to see available commands.",
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

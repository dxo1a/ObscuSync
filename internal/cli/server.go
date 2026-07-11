package cli

import (
	"github.com/spf13/cobra"

	"github.com/dxo1a/obscusync/internal/server"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start GameSync server",
	RunE: func(cmd *cobra.Command, args []string) error {
		return server.Start()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

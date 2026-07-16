package cli

import (
	"github.com/dxo1a/obscusync/internal/server"
	"github.com/spf13/cobra"
)

func (c *CLI) newServeCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Start GameSync server",
		RunE: func(
			cmd *cobra.Command,
			args []string,
		) error {
			return server.Start()
		},
	}
}

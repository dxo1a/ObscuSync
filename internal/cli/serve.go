package cli

import (
	"github.com/dxo1a/obscusync/internal/server"
	"github.com/spf13/cobra"
)

func (c *CLI) newServeCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Start GameSync HTTP server",
		RunE: func(cmd *cobra.Command, args []string) error {
			srv := server.New(
				c.config.ServerAddress(),
				c.storage,
				c.config,
			)

			return srv.Start()
		},
	}
}

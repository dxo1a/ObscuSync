package cli

import (
	"github.com/dxo1a/obscusync/internal/service"
	"github.com/spf13/cobra"
)

type CLI struct {
	root    *cobra.Command
	service *service.Service
}

func New(
	svc *service.Service,
) *CLI {
	cli := &CLI{
		service: svc,
	}

	cli.root = cli.newRootCommand()
	return cli
}

func (c *CLI) Execute() error {
	return c.root.Execute()
}

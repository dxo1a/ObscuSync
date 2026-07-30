package cli

import (
	"github.com/dxo1a/obscusync/internal/config"
	"github.com/dxo1a/obscusync/internal/manifest"
	"github.com/dxo1a/obscusync/internal/service"
	"github.com/spf13/cobra"
)

type CLI struct {
	root *cobra.Command

	service *service.Service

	config  *config.Manager
	storage *manifest.Storage
}

func New(svc *service.Service, cfg *config.Manager, storage *manifest.Storage) *CLI {
	cli := &CLI{
		service: svc,
		config:  cfg,
		storage: storage,
	}

	cli.root = cli.newRootCommand()
	return cli
}

func (c *CLI) Execute() error {
	return c.root.Execute()
}

package app

import (
	"github.com/dxo1a/obscusync/internal/cli"
	"github.com/dxo1a/obscusync/internal/config"
	"github.com/dxo1a/obscusync/internal/manifest"
	"github.com/dxo1a/obscusync/internal/scanner"
	"github.com/dxo1a/obscusync/internal/service"
)

func Run() {
	cfg, err := config.Load("configs/config.yaml")
	if err != nil {
		panic(err)
	}

	configManager := config.NewManager(cfg)
	sc := scanner.New()
	builder := manifest.New()
	storage := manifest.NewStorage(
		"data/manifests",
	)

	syncService := service.New(
		configManager,
		sc,
		builder,
		storage,
	)

	commandLine := cli.New(
		syncService,
		configManager,
		storage,
	)

	if err := commandLine.Execute(); err != nil {
		panic(err)
	}
}

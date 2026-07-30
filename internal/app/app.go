package app

import (
	"errors"
	"fmt"
	"os"

	"github.com/dxo1a/obscusync/internal/cli"
	"github.com/dxo1a/obscusync/internal/client"
	"github.com/dxo1a/obscusync/internal/config"
	"github.com/dxo1a/obscusync/internal/manifest"
	"github.com/dxo1a/obscusync/internal/scanner"
	"github.com/dxo1a/obscusync/internal/service"
)

func Run() {
	configPath, err := pathFromBase("config.yaml")
	if err != nil {
		fatal(err)
	}

	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		if err := os.WriteFile(configPath, []byte(defaultConfigTemplate), 0644); err != nil {
			fatal(fmt.Errorf("create config template: %w", err))
		}

		fmt.Printf(
			"Created config template:\n %s\n\nEdit this file (server/remote/profiles) and run the program again.\n",
			configPath,
		)
		os.Exit(0)
	}

	manifestsDir, err := pathFromBase("data", "manifests")
	if err != nil {
		fatal(err)
	}

	for _, parts := range [][]string{
		{"data", "cache"},
		{"data", "logs"},
		{"data", "manifests"},
	} {
		p, err := pathFromBase(parts...)
		if err != nil {
			fatal(err)
		}
		if err := os.MkdirAll(p, 0755); err != nil {
			fatal(err)
		}
	}

	cfg, err := config.Load(configPath)
	if err != nil {
		fatal(fmt.Errorf("load config %s: %w", configPath, err))
	}

	configManager := config.NewManager(cfg)
	sc := scanner.New()
	builder := manifest.New()
	storage := manifest.NewStorage(manifestsDir)
	httpClient := client.New(configManager.RemoteAddress())

	syncService := service.New(
		configManager,
		sc,
		builder,
		storage,
		httpClient,
	)

	commandLine := cli.New(
		syncService,
		configManager,
		storage,
	)

	if err := commandLine.Execute(); err != nil {
		fatal(err)
	}
}

func fatal(err error) {
	fmt.Fprint(os.Stderr, "Error: ", err)
	os.Exit(1)
}

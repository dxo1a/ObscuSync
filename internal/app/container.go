package app

import (
	"github.com/dxo1a/obscusync/internal/config"
	"github.com/dxo1a/obscusync/internal/manifest"
	"github.com/dxo1a/obscusync/internal/scanner"
	"github.com/dxo1a/obscusync/internal/service"
)

type Container struct {
	Config  *config.Manager
	Scanner *scanner.Scanner
	Builder *manifest.Builder
	Storage *manifest.Storage
	Service *service.Service
}

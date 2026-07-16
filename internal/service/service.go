package service

import (
	"path/filepath"

	"github.com/dxo1a/obscusync/internal/config"
	"github.com/dxo1a/obscusync/internal/manifest"
	"github.com/dxo1a/obscusync/internal/scanner"
)

type Service struct {
	config  *config.Manager
	scanner *scanner.Scanner
	builder *manifest.Builder
	storage *manifest.Storage
}

func New(
	cfg *config.Manager,
	scanner *scanner.Scanner,
	builder *manifest.Builder,
	storage *manifest.Storage,
) *Service {

	return &Service{
		config:  cfg,
		scanner: scanner,
		builder: builder,
		storage: storage,
	}
}

func (s *Service) Scan(profileName string) (*ScanResult, error) {

	profile, err := s.config.Profile(profileName)
	if err != nil {
		return nil, err
	}

	game, err := s.config.Game(profile)
	if err != nil {
		return nil, err
	}

	var scanFolders []string

	for _, folder := range game.ScanFolders {
		scanFolders = append(
			scanFolders,
			filepath.Join(profile.Root, folder),
		)
	}

	files, err := s.scanner.Scan(scanFolders)
	if err != nil {
		return nil, err
	}

	m := s.builder.Build(files)

	if err := s.storage.Save(profile.Name, m); err != nil {
		return nil, err
	}

	return &ScanResult{
		FileCount: len(files),
		Manifest: filepath.Join(
			"data",
			"manifests",
			profile.Name+".json",
		),
	}, nil
}

package service

import (
	"path/filepath"

	"github.com/dxo1a/obscusync/internal/config"
	"github.com/dxo1a/obscusync/internal/manifest"
	"github.com/dxo1a/obscusync/internal/scanner"
)

type Service struct {
	scanner *scanner.Scanner
	builder *manifest.Builder
	loader  *config.Loader
}

func New() *Service {
	return &Service{
		scanner: scanner.New(),
		builder: manifest.New(),
		loader:  config.NewLoader(),
	}
}

func (s *Service) Scan(profileName string) (*ScanResult, error) {
	err := s.loader.Load("configs/config.yaml")
	if err != nil {
		return nil, err
	}

	profile, err := s.loader.FindProfile(profileName)
	if err != nil {
		return nil, err
	}

	game, err := s.loader.FindGame(profile)
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

	err = manifest.Save("manifest.json", m)
	if err != nil {
		return nil, err
	}

	return &ScanResult{
		FileCount: len(files),
	}, nil
}

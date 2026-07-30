package service

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dxo1a/obscusync/internal/client"
	"github.com/dxo1a/obscusync/internal/config"
	"github.com/dxo1a/obscusync/internal/manifest"
	"github.com/dxo1a/obscusync/internal/scanner"
)

type Service struct {
	config  *config.Manager
	scanner *scanner.Scanner
	builder *manifest.Builder
	storage *manifest.Storage
	client  *client.Client
}

func New(
	cfg *config.Manager,
	scanner *scanner.Scanner,
	builder *manifest.Builder,
	storage *manifest.Storage,
	httpClient *client.Client,
) *Service {

	return &Service{
		config:  cfg,
		scanner: scanner,
		builder: builder,
		storage: storage,
		client:  httpClient,
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

func (s *Service) Sync(profileName string) (*SyncResult, error) {
	profile, err := s.config.Profile(profileName)
	if err != nil {
		return nil, err
	}

	game, err := s.config.Game(profile)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Fetching manifest from %s ...\n", s.client.BaseURL())

	remote, err := s.client.FetchManifest(profile.Name)
	if err != nil {
		return nil, err
	}

	// scan only existing folders (they may not be on the client yet)
	var scanFolders []string
	for _, folder := range game.ScanFolders {
		p := filepath.Join(profile.Root, folder)
		if _, err := os.Stat(p); err != nil {
			continue
		}
		scanFolders = append(scanFolders, p)
	}

	var localFiles []scanner.FileInfo
	if len(scanFolders) > 0 {
		localFiles, err = s.scanner.Scan(scanFolders)
		if err != nil {
			return nil, err
		}
	}

	localByPath := make(map[string]scanner.FileInfo, len(localFiles))
	for _, f := range localFiles {
		localByPath[f.Path] = f
	}

	remoteByPath := make(map[string]manifest.File, len(remote.Files))
	for _, f := range remote.Files {
		remoteByPath[f.Path] = f
	}

	result := &SyncResult{}

	for _, rf := range remote.Files {
		lf, exists := localByPath[rf.Path]
		if exists && lf.SHA256 == rf.SHA256 && lf.Size == rf.Size {
			result.Unchanged++
			continue
		}

		dest := filepath.Join(profile.Root, filepath.FromSlash(rf.Path))

		err := s.client.DownloadFile(profile.Name, rf.Path, dest, func(written, total int64) {
			if total <= 0 {
				total = rf.Size
			}
			printDownloadProgress(written, total, rf.Path)
		})
		if err != nil {
			finishDownloadProgress()
			return result, fmt.Errorf("download %s: %w", rf.Path, err)
		}
		finishDownloadProgress()
		result.Downloaded++
	}

	for _, lf := range localFiles {
		if _, ok := remoteByPath[lf.Path]; ok {
			continue
		}

		full := filepath.Join(profile.Root, filepath.FromSlash(lf.Path))
		if err := os.Remove(full); err != nil && !os.IsNotExist(err) {
			return result, fmt.Errorf("delete %s: %w", lf.Path, err)
		}
		fmt.Printf("deleted %s\n", lf.Path)
		result.Deleted++
	}

	return result, nil
}

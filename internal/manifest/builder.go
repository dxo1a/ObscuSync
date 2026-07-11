package manifest

import (
	"time"

	"github.com/dxo1a/obscusync/internal/scanner"
)

type Builder struct{}

func New() *Builder {
	return &Builder{}
}

func (b *Builder) Build(files []scanner.FileInfo) Manifest {
	manifest := Manifest{
		Version:   1,
		CreatedAt: time.Now().UTC(),
		Files:     make([]File, 0, len(files)),
	}

	for _, file := range files {
		manifest.Files = append(manifest.Files, File{
			Path:   file.Path,
			Size:   file.Size,
			SHA256: file.SHA256,
		})
	}

	return manifest
}

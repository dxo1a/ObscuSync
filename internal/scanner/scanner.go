package scanner

import (
	"io/fs"
	"path/filepath"
)

type Scanner struct {
}

func New() *Scanner {
	return &Scanner{}
}

func (s *Scanner) Scan(root string) ([]FileInfo, error) {
	var files []FileInfo

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {

		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		relativePath, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		info, err := d.Info()
		if err != nil {
			return err
		}

		hash, err := calculateSHA256(path)
		if err != nil {
			return err
		}

		files = append(files, FileInfo{
			Path:   filepath.ToSlash(relativePath),
			Size:   info.Size(),
			SHA256: hash,
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

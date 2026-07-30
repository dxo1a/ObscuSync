package app

import (
	"os"
	"path/filepath"
)

// baseDir - s the directory where the binary is located
// config.yaml and data/ are always relative to it not cwd
func baseDir() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}

	resolved, err := filepath.EvalSymlinks(exe)
	if err != nil {
		resolved = exe
	}

	return filepath.Dir(resolved), nil
}

func pathFromBase(parts ...string) (string, error) {
	base, err := baseDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(append([]string{base}, parts...)...), nil
}

package app

import (
	"os"
	"path/filepath"
)

// baseDir — каталог, в котором лежит бинарник.
// config.yaml и data/ всегда относительно него, а не cwd.
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

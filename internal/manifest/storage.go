package manifest

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Storage struct {
	root string
}

func NewStorage(root string) *Storage {
	return &Storage{
		root: root,
	}
}

func (s *Storage) Save(profile string, manifest Manifest) error {

	if err := os.MkdirAll(s.root, 0755); err != nil {
		return err
	}

	file, err := os.Create(
		filepath.Join(s.root, profile+".json"),
	)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	return encoder.Encode(manifest)
}

func (s *Storage) Load(profile string) (Manifest, error) {

	var manifest Manifest

	file, err := os.Open(
		filepath.Join(s.root, profile+".json"),
	)
	if err != nil {
		return manifest, err
	}

	defer file.Close()

	err = json.NewDecoder(file).Decode(&manifest)

	return manifest, err
}

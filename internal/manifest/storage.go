package manifest

import (
	"encoding/json"
	"os"
)

func Save(path string, manifest Manifest) error {
	data, err := json.MarshalIndent(manifest, "", "    ")
	if err != nil {
		return nil
	}

	return os.WriteFile(path, data, 0644)
}

func Load(path string) (Manifest, error) {
	var manifest Manifest

	data, err := os.ReadFile(path)
	if err != nil {
		return manifest, err
	}

	err = json.Unmarshal(data, &manifest)

	return manifest, err
}

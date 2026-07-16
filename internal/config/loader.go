package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

func Load(path string) (Config, error) {
	var cfg Config

	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}

	err = yaml.Unmarshal(data, &cfg)

	return cfg, err
}

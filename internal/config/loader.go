package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Loader struct {
	config Config
}

func NewLoader() *Loader {
	return &Loader{}
}

func (l *Loader) Load(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &l.config)
	if err != nil {
		return err
	}

	return nil
}

func (l *Loader) Config() Config {
	return l.config
}

func (l *Loader) FindProfile(name string) (Profile, error) {
	for _, profile := range l.config.Profiles {
		if profile.Name == name {
			return profile, nil
		}
	}
	return Profile{}, fmt.Errorf("profile '%s' not found", name)
}

func (l *Loader) FindGame(profile Profile) (Game, error) {
	game, ok := GetGame(profile.Game)
	if !ok {
		return Game{}, fmt.Errorf(
			"game '%s' is not supported",
			profile.Game,
		)
	}

	return game, nil
}

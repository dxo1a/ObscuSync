package config

type Manager struct {
	config Config
}

func NewManager(cfg Config) *Manager {
	return &Manager{
		config: cfg,
	}
}

func (m *Manager) Profile(name string) (Profile, error) {
	for _, profile := range m.config.Profiles {
		if profile.Name == name {
			return profile, nil
		}
	}

	return Profile{}, ErrProfileNotFound(name)
}

func (m *Manager) Game(profile Profile) (Game, error) {
	game, ok := GetGame(profile.Game)
	if !ok {
		return Game{}, ErrGameNotSupported(profile.Game)
	}

	return game, nil
}

func (m *Manager) ServerAddress() string {
	return m.config.Server.Address
}

func (m *Manager) RemoteAddress() string {
	return m.config.Remote.Address
}

package config

type Config struct {
	Server   ServerConfig `yaml:"server"`
	Remote   RemoteConfig `yaml: "remote"`
	Profiles []Profile    `yaml:"profiles"`
}

type ServerConfig struct {
	Address string `yaml:"address"`
}

type RemoteConfig struct {
	Address string `yaml:"address`
}

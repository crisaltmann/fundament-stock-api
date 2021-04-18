package config

type Config struct {
	ApplicationConfig `yaml:"application"`
}

type ApplicationConfig struct {
	Address string `yaml:"address"`
}


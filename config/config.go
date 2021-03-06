package config

type Config struct {
	ApplicationConfig `yaml:"application"`
	DatabaseConfig    `yaml:"database"`
	Job				  `yaml:"job"`
}

type ApplicationConfig struct {
	Address string `yaml:"address"`
}

type DatabaseConfig struct {
	Database string `yaml:"database"`
	Scheme   string `yaml:"scheme"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Hostname string `yaml:"hostname"`
	Port     int    `yaml:"port"`
	Url      string `yaml:"url"`
}

type Job struct {
	AssetSync `yaml:"asset-sync"`
}

type AssetSync struct {
	Cron	string `yaml:"cron"`
}
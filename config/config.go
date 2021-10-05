package config

type SourceInfo struct {}

type Config struct {
	sources []SourceInfo
}

func LoadConfig(configPath string) (*Config, error) {
	return new(Config), nil
}

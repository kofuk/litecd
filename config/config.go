package config

type SourceInfo struct {}

type Config struct {
	sources []SourceInfo
}

func LoadConfig() (*Config, error) {
	return new(Config), nil
}

package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type SourceInfo struct {
	FetcherType string `yaml:"type"`
	GitRepo     string `yaml:"githubRepo"`
}

type Config struct {
	Sources []SourceInfo `yaml:"sources"`
}

func LoadConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)

	return &config, nil
}

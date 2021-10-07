package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type Credential struct {
	CredentialType string `yaml:"type"`

	// Fields for CredentialType="password"
	UserName       string `yaml:"user"`
	Password       string `yaml:"password"`
}

type SourceInfo struct {
	FetcherType string     `yaml:"type"`

	// Fields for FetcherType="git"
	GitRepo     string     `yaml:"gitRepo"`

	Credential  Credential `yaml:"credential"`
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

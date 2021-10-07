package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type Credential struct {
	CredentialType string `yaml:"type"`

	// Fields for CredentialType="password"
	UserName string `yaml:"user"`
	Password string `yaml:"password"`
}

type SourceInfo struct {
	FetcherType string `yaml:"type"`

	// Fields for FetcherType="git"
	GitRepo   string `yaml:"gitRepo"`
	GitBranch string `yaml:"gitBranch"`

	Credential Credential `yaml:"credential"`
}

type Config struct {
	Sources []SourceInfo `yaml:"sources"`
}

func LoadConfig(fs Filesystem) (*Config, error) {
	file, err := os.Open(fs.GetConfigPath())
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

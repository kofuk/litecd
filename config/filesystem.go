package config

import (
	"os"
	"path/filepath"
)

type Filesystem struct {
	dataRoot string
}

const (
	configPath = "/etc/litecd.yml"
	dataDir    = "/var/lib/litecd"
)

func FilesystemNew() Filesystem {
	return Filesystem{
		dataRoot: os.Getenv("LITECD_DATA_ROOT"),
	}
}

func (h Filesystem) GetConfigPath() string {
	return filepath.Join(h.dataRoot, configPath)
}

func (h Filesystem) PrepareDataDir() (string, error) {
	path := filepath.Join(h.dataRoot, dataDir)
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err)  {
		err := os.MkdirAll(path, 0700)
		if err != nil {
			return "", err
		}
	} else if err != nil {
		return "", err
	}

	return path, nil
}

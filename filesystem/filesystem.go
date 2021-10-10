package filesystem

import (
	"os"
	"path/filepath"
)

type DiskFilesystem struct {
	dataRoot string
}

type Filesystem interface {
	GetConfigPath() string
	PrepareDataDir() (string, error)
	PrepareSecretsDir() (string, error)
}

const (
	configPath  = "/etc/litecd.yml"
	dataBaseDir = "/var/lib/litecd"
	secretsDir  = "secrets"
)

func FilesystemNew() Filesystem {
	return DiskFilesystem{
		dataRoot: os.Getenv("LITECD_DATA_ROOT"),
	}
}

func (fs DiskFilesystem) GetConfigPath() string {
	return filepath.Join(fs.dataRoot, configPath)
}

func (fs DiskFilesystem) PrepareDataDir() (string, error) {
	path := filepath.Join(fs.dataRoot, dataBaseDir)
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return "", err
		}
	} else if err != nil {
		return "", err
	}

	return path, nil
}

func (fs DiskFilesystem) PrepareSecretsDir() (string, error) {
	if _, err := fs.PrepareDataDir(); err != nil {
		return "", err
	}

	path := filepath.Join(fs.dataRoot, dataBaseDir, secretsDir)
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		if err := os.Mkdir(path, 0700); err != nil {
			return "", err
		}
	} else if err != nil {
		return "", err
	}
	return path, nil
}

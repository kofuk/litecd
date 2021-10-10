package secrets

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/kofuk/litecd/config"
)

const (
	secretsFilename = "secrest.json"
)

type Secrets map[string]string

func GetAllSecrets(fs config.Filesystem) (Secrets, error) {
	secretsDir, err := fs.PrepareSecretsDir()
	if err != nil {
		return nil, err
	}

	secretsPath := filepath.Join(secretsDir, secretsFilename)

	file, err := os.Open(secretsPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	result := make(Secrets)
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func StoreSecret(fs config.Filesystem, key, value string) error {
	secretsData, err := GetAllSecrets(fs)
	if err != nil {
		return err
	}

	secretsData[key] = value

	secretsDir, err := fs.PrepareSecretsDir()
	if err != nil {
		return err
	}

	secretsPath := filepath.Join(secretsDir, secretsFilename)

	data, err := json.Marshal(secretsData)

	file, err := os.Create(secretsPath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.Write(data)
	if err != nil {
		return err
	}
	writer.Flush()

	return nil
}

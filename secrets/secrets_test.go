package secrets

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/kofuk/litecd/config"
)

func setUpWithFile(name, testJson string) (string, error) {
	dataRoot := filepath.Join("/tmp/litecd_test.d", name)
	secretsPath := filepath.Join(dataRoot, "var/lib/litecd/secrets")

	err := os.MkdirAll(secretsPath, 0755)
	if err != nil && !os.IsExist(err) {
		return "", err
	}

	infile, err := os.Open(testJson)
	if err != nil {
		return "", err
	}
	defer infile.Close()

	outfile, err := os.Create(filepath.Join(secretsPath, secretsFilename))
	if err != nil {
		return "", err
	}
	defer outfile.Close()

	_, err = io.Copy(outfile, infile)
	if err != nil {
		return "", err
	}

	return dataRoot, nil
}

func tearDown(name string) error {
	dataRoot := filepath.Join("/tmp/litecd_test.d", name)
	if err := os.RemoveAll(dataRoot); !os.IsNotExist(err) {
		return err
	}
	return nil
}

func TestGetAllSecrets(t *testing.T) {
	path, err := setUpWithFile("read", "testdata/test_secrets.json")
	if err != nil {
		t.Fatal("Setup failure: ", err)
	}
	defer tearDown("read")

	os.Setenv("LITECD_DATA_ROOT", path)
	fs := config.FilesystemNew()

	secrets, err := GetAllSecrets(fs)
	if err != nil {
		t.Fatal(err)
	}

	if secrets["foo"] != "bar" {
		t.Fatalf("expects secrets[\"foo\"]==bar, but got %s\n", secrets["foo"])
	}
}

func TestStoreSecret(t *testing.T) {
	path, err := setUpWithFile("write", "testdata/test_secrets.json")
	if err != nil {
		t.Fatal("Setup failure: ", err)
	}
	defer tearDown("write")

	os.Setenv("LITECD_DATA_ROOT", path)
	fs := config.FilesystemNew()

	err = StoreSecret(fs, "hoge", "fuga")
	if err != nil {
		t.Fatal(err)
	}

	secrets, err := GetAllSecrets(fs)
	if err != nil {
		t.Fatal(err)
	}

	if secrets["foo"] != "bar" {
		t.Fatalf("expects secrets[\"foo\"]==bar, but got %s\n", secrets["foo"])
	}

	if secrets["hoge"] != "fuga" {
		t.Fatalf("expects secrets[\"hoge\"]==fuga, but got %s\n", secrets["hoge"])
	}
}

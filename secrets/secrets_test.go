package secrets

import (
	"io"
	"os"
	"path/filepath"
	"testing"
)

type fakeFilesystem struct {
	rootDir string
}

func newFakeFilesystem(name string) fakeFilesystem {
	rootDir := filepath.Join("/tmp/litecd_test.d", name)
	if err := os.MkdirAll(rootDir, 0755); err != nil {
		panic(err)
	}

	// Prepare secret file
	infile, err := os.Open(filepath.Join("testdata/", name+".json"))
	if err != nil {
		panic(err)
	}
	defer infile.Close()

	outfile, err := os.Create(filepath.Join(rootDir, "secrets.json"))
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	if _, err = io.Copy(outfile, infile); err != nil {
		panic(err)
	}

	return fakeFilesystem{
		rootDir: rootDir,
	}
}

func (fs fakeFilesystem) GetConfigPath() string {
	return filepath.Join(fs.rootDir, "config.yml")
}

func (fs fakeFilesystem) PrepareDataDir() (string, error) {
	return fs.rootDir, nil
}

func (fs fakeFilesystem) PrepareSecretsDir() (string, error) {
	return fs.rootDir, nil
}

func (fs fakeFilesystem) tearDown() {
	if err := os.RemoveAll(fs.rootDir); err != nil {
		panic(err)
	}
}

func TestGetAllSecrets(t *testing.T) {
	fs := newFakeFilesystem("test_secrets")
	defer fs.tearDown()

	secrets, err := GetAllSecrets(fs)
	if err != nil {
		t.Fatal(err)
	}

	if secrets["foo"] != "bar" {
		t.Fatalf("expects secrets[\"foo\"]==bar, but got %s\n", secrets["foo"])
	}
}

func TestStoreSecret(t *testing.T) {
	fs := newFakeFilesystem("test_secrets")
	defer fs.tearDown()

	err := StoreSecret(fs, "hoge", "fuga")
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

package fetcher

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/kofuk/litecd/config"
	"github.com/kofuk/litecd/filesystem"
)

type GitFetcher struct{}

const (
	baseDir = "git"
)

func init() {
	fetchers["git"] = &GitFetcher{}
}

func createDataDirName(source *config.SourceInfo) string {
	var builder strings.Builder
	builder.WriteString(source.GitRepo)
	builder.WriteRune('[')
	builder.WriteString(source.GitBranch)
	builder.WriteRune(']')

	data := sha256.Sum256([]byte(builder.String()))
	return string(hex.EncodeToString(data[:]))
}

func getDataPath(dataDir string, source *config.SourceInfo) string {
	return filepath.Join(dataDir, "git", createDataDirName(source))
}

func (*GitFetcher) IsInitialized(source *config.SourceInfo, fs filesystem.Filesystem) (bool, error) {
	path, err := fs.PrepareDataDir()
	if err != nil {
		return false, err
	}

	gitDir := filepath.Join(getDataPath(path, source), ".git")
	_, err = os.Stat(gitDir)
	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func (*GitFetcher) Initialize(source *config.SourceInfo, fs filesystem.Filesystem) error {
	path, err := fs.PrepareDataDir()
	if err != nil {
		return err
	}

	base := filepath.Join(path, baseDir)
	if _, err := os.Stat(base); err != nil && os.IsNotExist(err) {
		if err := os.MkdirAll(base, 0755); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	dest := filepath.Join(base, createDataDirName(source))

	args := []string{"clone"}
	if source.GitBranch != "" {
		args = append(args, fmt.Sprintf("--branch=%s", source.GitBranch))
	}
	args = append(args, "--", source.GitRepo, dest)

	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (*GitFetcher) Update(source *config.SourceInfo, fs filesystem.Filesystem) error {
	dataDir, err := fs.PrepareDataDir()
	if err != nil {
		return err
	}

	repo := getDataPath(dataDir, source)

	cmd := exec.Command("git", "-C", repo, "fetch")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("git", "-C", repo, "reset", "--hard", "FETCH_HEAD")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

package fetcher

import (
	"errors"

	"github.com/kofuk/litecd/config"
	"github.com/kofuk/litecd/filesystem"
)

type Fetcher interface {
	IsInitialized(*config.SourceInfo, filesystem.Filesystem) (bool, error)
	Initialize(*config.SourceInfo, filesystem.Filesystem) error
	Update(*config.SourceInfo, filesystem.Filesystem) error
}

var (
	NoSuchFetcher = errors.New("No such fetcher")
)

var fetchers = make(map[string]Fetcher)

func GetFetcherForType(sourceType string) (Fetcher, error) {
	fetcher := fetchers[sourceType]
	if fetcher == nil {
		return nil, NoSuchFetcher
	}
	return fetcher, nil
}

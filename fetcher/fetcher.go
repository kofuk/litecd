package fetcher

import (
	"errors"

	"github.com/kofuk/litecd/config"
)

type Fetcher interface {
	IsInitialized(*config.SourceInfo, config.Filesystem) (bool, error)
	Initialize(*config.SourceInfo, config.Filesystem) error
	Update(*config.SourceInfo, config.Filesystem) error
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

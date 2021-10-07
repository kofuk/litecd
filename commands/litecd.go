package commands

import (
	"log"

	"github.com/kofuk/litecd/config"
	"github.com/kofuk/litecd/fetcher"
	"github.com/spf13/cobra"
)

const (
	configFile = "/etc/litecd.yml"
)

var rootCmd = cobra.Command{
	Use:   "litecd",
	Short: "Tiny deployment infrastructure focused on personal projects.",
	Long:  "",
	RunE:  runApp,
}

func runApp(cmd *cobra.Command, args []string) error {
	fs := config.FilesystemNew()

	config, err := config.LoadConfig(fs)
	if err != nil {
		return err
	}

	for _, source := range config.Sources {
		fetcher, err := fetcher.GetFetcherForType(source.FetcherType)
		if err != nil {
			return err
		}

		if initialized, _ := fetcher.IsInitialized(&source, fs); !initialized {
			if err := fetcher.Initialize(&source, fs); err != nil {
				log.Println(err)
			}
		} else {
			if err := fetcher.Update(&source, fs); err != nil {
				log.Println(err)
			}
		}
	}

	return nil
}

func Execute() error {
	return rootCmd.Execute()
}

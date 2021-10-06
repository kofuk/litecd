package commands

import (
	"path/filepath"

	"github.com/kofuk/litecd/config"
	"github.com/spf13/cobra"
)

const (
	configFile = "etc/litecd.yml"
)

type AppOption struct {
	dataRoot string
}

var option AppOption

func init() {
	rootCmd.PersistentFlags().StringVar(&option.dataRoot, "data-root", "/",
		"Data directory for litecd (default: /)")
}

var rootCmd = cobra.Command{
	Use:   "litecd",
	Short: "Tiny deployment infrastructure focused on personal projects.",
	Long:  "",
	RunE:  runApp,
}

func runApp(cmd *cobra.Command, args []string) error {
	config, err := config.LoadConfig(filepath.Join(option.dataRoot, configFile))
	if err != nil {
		return err
	}

	_ = config

	return nil
}

func Execute() error {
	return rootCmd.Execute()
}

package commands

import (
	"os"
	"path/filepath"

	"github.com/kofuk/litecd/config"
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
	dataRoot := os.Getenv("LITECD_DATA_ROOT")

	config, err := config.LoadConfig(filepath.Join(dataRoot, configFile))
	if err != nil {
		return err
	}

	_ = config

	return nil
}

func Execute() error {
	return rootCmd.Execute()
}

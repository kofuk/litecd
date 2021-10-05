package main

import (
	"fmt"
	"os"

	"github.com/kofuk/litecd/config"
	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:   "litecd",
	Short: "Tiny deployment infrastructure focused on personal projects.",
	Long:  "",
	RunE:  runApp,
}

type AppConfig struct {
	configFilePath string
}

var appConfig AppConfig

func init() {
	rootCmd.PersistentFlags().StringVar(&appConfig.configFilePath, "config", "/etc/litecd.conf",
		"Configuration file (default: /etc/litecd.conf)")
}

func runApp(cmd *cobra.Command, args []string) error {
	config, err := config.LoadConfig(appConfig.configFilePath)
	if err != nil {
		return err
	}

	_ = config

	return nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

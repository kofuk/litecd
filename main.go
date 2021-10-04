package main

import (
	"fmt"
	"os"

	"github.com/kofuk/litecd/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading configuration: %v\n", err)
	}

	_ = config
}

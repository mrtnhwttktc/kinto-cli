package main

import (
	"os"

	"github.com/mrtnhwttktc/kinto-cli/cmd/cli"
	"github.com/mrtnhwttktc/kinto-cli/internal/config"
	"github.com/mrtnhwttktc/kinto-cli/internal/logger"
)

func main() {
	// Set the default logger
	logger.SetDefaultLogger()
	// Load the configuration
	config.InitializeConfig()

	// Execute the root command to initialize the CLI
	rootCmd := cli.NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

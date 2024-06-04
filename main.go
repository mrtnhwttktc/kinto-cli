package main

import (
	"os"

	"github.com/mrtnhwttktc/kinto-cli/cmd/cli"
	"github.com/mrtnhwttktc/kinto-cli/internal/config"
	"github.com/mrtnhwttktc/kinto-cli/internal/localizer"
	"github.com/mrtnhwttktc/kinto-cli/internal/logger"
)

func main() {
	// Set the default logger
	logger.SetDefaultLogger()
	// Load the configuration
	_ = config.GetConfig()
	// Load the localizer
	_ = localizer.GetLocalizer()

	// Execute the root command to initialize the CLI
	rootCmd := cli.NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

package main

import (
	"os"

	"github.com/mrtnhwttktc/kinto-cli/cmd"
	"github.com/mrtnhwttktc/kinto-cli/internal/config"
	"github.com/mrtnhwttktc/kinto-cli/internal/localizer"
	_ "github.com/mrtnhwttktc/kinto-cli/internal/logger"
)

func main() {
	// Load the configuration
	_ = config.GetConfig()
	_ = localizer.GetLocalizer()
	rootCmd := cmd.NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

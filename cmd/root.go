package cmd

import (
	"log/slog"
	"os"

	"github.com/mrtnhwttktc/kinto-cli/internal/localizer"
	ktcliLogger "github.com/mrtnhwttktc/kinto-cli/internal/logger"
	v "github.com/mrtnhwttktc/kinto-cli/internal/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "ktcli",
	Version: v.Version,
	Short:   localizer.Lang.Translate("CLI for Kinto scripts and tools."),
	Long:    localizer.Lang.Translate(`Kinto CLI or ktcli is a command line interface for employees at Kinto Technologies, allowing easy access to the multiple tools and scripts developped by our teams.`),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		debug, _ := cmd.Flags().GetBool("debug")
		verbose, _ := cmd.Flags().GetBool("verbose")
		if verbose {
			ktcliLogger.VerboseLogger()
			slog.Info("Verbose mode enabled.")
		}
		if debug {
			ktcliLogger.LogLevel.Set(slog.LevelDebug)
			slog.Info("Log level set to debug.")
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().Bool("debug", false, localizer.Lang.Translate("Sets the log level to debug."))
	rootCmd.PersistentFlags().Bool("verbose", false, localizer.Lang.Translate("Prints logs to stdout."))
}

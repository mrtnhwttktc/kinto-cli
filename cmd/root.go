package cmd

import (
	"log/slog"

	"github.com/mrtnhwttktc/kinto-cli/cmd/set"
	"github.com/mrtnhwttktc/kinto-cli/internal/localizer"
	ktcliLogger "github.com/mrtnhwttktc/kinto-cli/internal/logger"
	v "github.com/mrtnhwttktc/kinto-cli/internal/version"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	l := localizer.GetLocalizer()
	rootCmd := &cobra.Command{
		Use:     "ktcli",
		Version: v.Version,
		Short:   l.Translate("CLI for Kinto scripts and tools."),
		Long:    l.Translate(`Kinto CLI or ktcli is a command line interface for employees at Kinto Technologies, allowing easy access to the multiple tools and scripts developped by our teams.`),
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

	bindFlags(rootCmd, l)

	rootCmd.AddCommand(
		set.NewSetCmd(),
	)
	return rootCmd
}

func bindFlags(cmd *cobra.Command, l *localizer.Localizer) {
	cmd.PersistentFlags().Bool("debug", false, l.Translate("Sets the log level to debug."))
	cmd.PersistentFlags().Bool("verbose", false, l.Translate("Prints logs to stdout."))
}

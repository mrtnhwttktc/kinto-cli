package cli

import (
	"log/slog"

	cc "github.com/ivanpirog/coloredcobra"
	"github.com/mrtnhwttktc/kinto-cli/cmd/cli/set"
	"github.com/mrtnhwttktc/kinto-cli/cmd/utils"
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

	// Disable the default completion command
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	bindFlags(rootCmd, l)

	rootCmd.AddCommand(
		set.NewSetCmd(),
	)

	// colorize the help output. Needs to be called before LocalizeUsageTemplate
	cc.Init(&cc.Config{
		RootCmd:         rootCmd,
		Headings:        cc.HiCyan + cc.Bold,
		Commands:        cc.HiYellow + cc.Bold,
		CmdShortDescr:   cc.Blue,
		Example:         cc.Italic,
		ExecName:        cc.Bold,
		Flags:           cc.Bold,
		FlagsDescr:      cc.Blue,
		FlagsDataType:   cc.Italic,
		NoExtraNewlines: true,
		NoBottomNewline: true,
	})

	utils.LocalizeHelpFunc(rootCmd, l)
	utils.LocalizeUsageTemplate(rootCmd, l)

	return rootCmd
}

func bindFlags(cmd *cobra.Command, l *localizer.Localizer) {
	cmd.PersistentFlags().Bool("debug", false, l.Translate("Sets the log level to debug."))
	cmd.PersistentFlags().Bool("verbose", false, l.Translate("Prints logs to stdout."))
	cmd.Flags().BoolP("version", "v", false, l.Translate("Prints the version of the CLI."))
	cmd.Flags().BoolP("help", "h", false, l.Translate("Prints the help message."))
}

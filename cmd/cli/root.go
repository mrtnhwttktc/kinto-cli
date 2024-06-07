package cli

import (
	"fmt"
	"io"
	"log/slog"

	cc "github.com/ivanpirog/coloredcobra"
	"github.com/mrtnhwttktc/kinto-cli/cmd/cli/set"
	"github.com/mrtnhwttktc/kinto-cli/cmd/cli/update"
	"github.com/mrtnhwttktc/kinto-cli/cmd/utils"
	"github.com/mrtnhwttktc/kinto-cli/internal/localizer"
	"github.com/mrtnhwttktc/kinto-cli/internal/logger"
	v "github.com/mrtnhwttktc/kinto-cli/internal/version"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func NewRootCmd() *cobra.Command {
	l := localizer.NewLocalizer()

	rootCmd := &cobra.Command{
		Use:     "ktcli",
		Version: v.Version,
		Short:   l.Translate("CLI for Kinto scripts and tools."),
		Long:    l.Translate(`Kinto CLI or ktcli is a command line interface for employees at Kinto Technologies, allowing easy access to the multiple tools and scripts developped by our teams.`),
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Bind the flags to the config, so that the config file and environment variables are used if the flag is not set
			bindFlagsToEnvAndConfig(cmd)
			logger.SetLogLevel(viper.GetString("log-level"))

			debug, _ := cmd.Flags().GetBool("debug")
			verbose, _ := cmd.Flags().GetBool("verbose")
			quiet, _ := cmd.Flags().GetBool("quiet")
			no_version_check, _ := cmd.Flags().GetBool("no-version-check")

			if !no_version_check {
				if v.IsNewVersionAvailable() {
					out := cmd.OutOrStdout()
					fmt.Fprintf(out, "\033[35m%s\n%s:\033[0m \033[32msudo ktcli update\n\n\033[0m", l.Translate("A new version of the CLI is available. Please update to the latest version."), l.Translate("Use the update command to update"))
				}
			}
			if verbose {
				logger.VerboseLogger()
				slog.Info("Verbose mode enabled.")
			}
			if debug {
				logger.LogLevel.Set(slog.LevelDebug)
				slog.Info("Log level set to debug.")
			}
			if quiet {
				cmd.SetOut(io.Discard)
			}
		},
	}

	setFlags(rootCmd, l)

	// Disable the default completion command
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// Add commands here
	rootCmd.AddCommand(
		set.NewSetCmd(),
		update.NewUpdateCmd(),
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

	// Localize the help function, usage template and version template
	utils.LocalizeHelpFunc(rootCmd, l)
	utils.LocalizeUsageTemplate(rootCmd, l)
	utils.LocalizeVersionTemplate(rootCmd, l)

	return rootCmd
}

func setFlags(cmd *cobra.Command, l *localizer.Localizer) {
	// global flags
	cmd.PersistentFlags().BoolP("help", "h", false, l.Translate("help for %s", cmd.Name()))
	cmd.PersistentFlags().Bool("debug", false, l.Translate("Sets the log level to debug."))
	cmd.PersistentFlags().Bool("verbose", false, l.Translate("Prints logs to stdout."))
	cmd.PersistentFlags().Bool("no-version-check", false, l.Translate("Disables the check for a new version of the CLI."))
	cmd.PersistentFlags().BoolP("non-interactive", "n", false, l.Translate("Disables interactive mode."))
	cmd.PersistentFlags().BoolP("quiet", "q", false, l.Translate("Disables all output except errors."))
	cmd.PersistentFlags().String("log-level", "info", l.Translate("Sets the log level. Options: debug, info, warn, error."))

	// local flags
	cmd.Flags().BoolP("version", "v", false, l.Translate("Prints the version of the CLI."))

	// hide the debug, verbose and no-version-check flags from the help output
	cmd.PersistentFlags().MarkHidden("debug")
	cmd.PersistentFlags().MarkHidden("verbose")
	cmd.PersistentFlags().MarkHidden("no-version-check")
}

// Bind each cobra flag to its associated viper configuration (config file and environment variable)
// The order of precedence is default value < config file < environment variable < flag
func bindFlagsToEnvAndConfig(cmd *cobra.Command) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		configName := f.Name
		if !f.Changed && viper.IsSet(configName) {
			val := viper.Get(configName)
			cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}

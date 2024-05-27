package cli

import (
	"fmt"
	"log/slog"

	cc "github.com/ivanpirog/coloredcobra"
	"github.com/mrtnhwttktc/kinto-cli/cmd/cli/set"
	"github.com/mrtnhwttktc/kinto-cli/cmd/cli/update"
	"github.com/mrtnhwttktc/kinto-cli/cmd/utils"
	"github.com/mrtnhwttktc/kinto-cli/internal/config"
	"github.com/mrtnhwttktc/kinto-cli/internal/localizer"
	ktcliLogger "github.com/mrtnhwttktc/kinto-cli/internal/logger"
	v "github.com/mrtnhwttktc/kinto-cli/internal/version"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func NewRootCmd() *cobra.Command {
	l := localizer.GetLocalizer()

	rootCmd := &cobra.Command{
		Use:     "ktcli",
		Version: v.Version,
		Short:   l.Translate("CLI for Kinto scripts and tools."),
		Long:    l.Translate(`Kinto CLI or ktcli is a command line interface for employees at Kinto Technologies, allowing easy access to the multiple tools and scripts developped by our teams.`),
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Bind the flags to the config, so that the config file and environment variables are used if the flag is not set
			bindFlagsToEnvAndConfig(cmd, config.GetConfig())

			debug, _ := cmd.Flags().GetBool("debug")
			verbose, err := cmd.Flags().GetBool("verbose")
			if err != nil {
				slog.Error("Error getting verbose flag.", slog.String("error", err.Error()))
			}
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

	// Localize the help function and usage template
	utils.LocalizeHelpFunc(rootCmd, l)
	utils.LocalizeUsageTemplate(rootCmd, l)

	return rootCmd
}

func setFlags(cmd *cobra.Command, l *localizer.Localizer) {
	// global flags
	cmd.PersistentFlags().BoolP("help", "h", false, l.Translate("help for %s", cmd.Name()))
	cmd.PersistentFlags().Bool("debug", false, l.Translate("Sets the log level to debug."))
	cmd.PersistentFlags().Bool("verbose", false, l.Translate("Prints logs to stdout."))
	cmd.PersistentFlags().BoolP("non-interactive", "n", false, l.Translate("Disables interactive mode."))
	// local flags
	cmd.Flags().BoolP("version", "v", false, l.Translate("Prints the version of the CLI."))
}

// Bind each cobra flag to its associated viper configuration (config file and environment variable)
// The order of precedence is flag > environment variable > config file > default value
func bindFlagsToEnvAndConfig(cmd *cobra.Command, conf *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		configName := f.Name
		if !f.Changed && conf.IsSet(configName) {
			val := conf.Get(configName)
			cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}

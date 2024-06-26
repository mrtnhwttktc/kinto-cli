package set

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/mrtnhwttktc/kinto-cli/cmd/cli/set/language"
	"github.com/mrtnhwttktc/kinto-cli/internal/localizer"
	"github.com/spf13/cobra"
)

func NewSetCmd() *cobra.Command {
	l := localizer.NewLocalizer()
	setCmd := &cobra.Command{
		Use:   "set",
		Short: l.Translate("Set global configurations."),
		Long:  l.Translate("Set global configurations for ktcli. Updates the local configuration file with the selected option.\nIn interactive mode, the CLI will prompt you to select a configuration to set. In non-interactive mode, you must provide a subcommand."),
		Example: `
		# interactive mode
		ktcli set
		
		# non-interactive mode
		ktcli set language english
		`,
		PreRun: func(cmd *cobra.Command, args []string) {
			n, _ := cmd.Flags().GetBool("non-interactive")
			if n {
				out := cmd.OutOrStdout()
				slog.Warn("Cannot use non-interactive mode with set command. A subcommand is required.")
				fmt.Fprintln(out, l.Translate("Cannot use non-interactive mode with set command. A subcommand is required.\n"))
				cmd.Help()
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			out := cmd.OutOrStdout()
			var subcommands = []string{}
			for _, subcmd := range cmd.Commands() {
				subcommands = append(subcommands, subcmd.Name())
			}
			prompt := promptui.Select{
				Label: l.Translate("Select the configuration to set"),
				Items: subcommands,
			}
			_, result, err := prompt.Run()
			if err != nil {
				slog.Error("Subcommand selection prompt failed", slog.String("error", err.Error()))
				return
			}
			fmt.Fprintln(out, l.Translate("Running subcommand %s", result))
			for _, subcmd := range cmd.Commands() {
				if subcmd.Name() == result {
					subcmd.Run(cmd, []string{})
				}
			}
		},
	}
	setFlags(setCmd, l)

	setCmd.AddCommand(language.NewLanguageCmd())

	return setCmd
}

func setFlags(cmd *cobra.Command, l *localizer.Localizer) {
}

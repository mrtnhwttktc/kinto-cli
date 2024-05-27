package language

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/mrtnhwttktc/kinto-cli/internal/config"
	"github.com/mrtnhwttktc/kinto-cli/internal/localizer"
	"github.com/spf13/cobra"
)

func NewLanguageCmd() *cobra.Command {
	l := localizer.GetLocalizer()
	languageCmd := &cobra.Command{
		Use:   "language [english|japanese]",
		Short: l.Translate("Set the language to use for the CLI."),
		Long:  l.Translate("Set the language to use for the CLI. Updates the local configuration file with the selected language. If no language is provided, the CLI will prompt you to select one."),
		Example: `
	# interactive mode
	ktcli set language

	# non-interactive mode
	ktcli set language english
	`,
		PreRun: func(cmd *cobra.Command, args []string) {
			n, _ := cmd.Flags().GetBool("non-interactive")
			if n && len(args) == 0 {
				slog.Error("Cannot use non-interactive mode with set language command and no arguments. A language must be provided as an argument.")
				fmt.Println(l.Translate("Cannot use non-interactive mode with set language command and no arguments. A language must be provided as an argument.\n"))
				cmd.Help()
				os.Exit(1)
			}
			if len(args) > 1 {
				slog.Error("Too many arguments provided for set language command.")
				fmt.Println(l.Translate("Too many arguments provided for set language command. Please select a language from the following: %v\n", localizer.GetLangOptions()))
				cmd.Help()
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			langs := localizer.GetLangOptions()
			conf := config.GetConfig()
			if len(args) == 1 {
				fmt.Println(l.Translate("You selected: %s", args[0]))
				err := config.SaveToConfig(conf, "language", args[0])
				if err != nil {
					slog.Error("Error saving language to config.", slog.String("error", err.Error()))
					fmt.Println(l.Translate("Error saving language to config."))
					os.Exit(1)
				}
				return
			}

			prompt := promptui.Select{
				Label: l.Translate("Select the language to use"),
				Items: langs,
			}

			_, result, err := prompt.Run()
			if err != nil {
				slog.Error("Language selection prompt failed", slog.String("error", err.Error()))
				return
			}
			fmt.Println(l.Translate("You selected: %s", result))
			err = config.SaveToConfig(conf, "language", result)
			if err != nil {
				slog.Error("Error saving language to config.", slog.String("error", err.Error()))
				fmt.Println(l.Translate("Error saving language to config."))
				os.Exit(1)
			}
		},
	}
	setFlags(languageCmd, l)
	return languageCmd
}

func setFlags(cmd *cobra.Command, l *localizer.Localizer) {
}

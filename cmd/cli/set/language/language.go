package language

import (
	"fmt"
	"io"
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
			out := cmd.OutOrStdout()
			n, _ := cmd.Flags().GetBool("non-interactive")
			if n && len(args) == 0 {
				fmt.Fprintln(out, l.Translate("Cannot use non-interactive mode with set language command and no arguments. A language must be provided as an argument.\n"))
				cmd.Help()
				os.Exit(1)
			}
			if len(args) > 1 {
				fmt.Fprintln(out, l.Translate("Too many arguments provided for set language command. Please select a language from the following: %v\n", localizer.GetLangOptions()))
				cmd.Help()
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			out := cmd.OutOrStdout()
			langs := localizer.GetLangOptions()
			if len(args) == 1 {
				for _, lang := range langs {
					if lang == args[0] {
						setLanguage(out, l, args[0])
						return
					}
				}
				fmt.Fprintln(out, l.Translate("Invalid language %s provided. Valid languages: %v\n", args[0], langs))
				n, _ := cmd.Flags().GetBool("non-interactive")
				if n {
					os.Exit(1)
				}
				fmt.Fprintln(out, l.Translate("Please select a language."))
			}

			result, err := interactiveMode(l, langs)
			if err != nil {
				return
			}
			setLanguage(out, l, result)
		},
	}
	setFlags(languageCmd, l)
	return languageCmd
}

func interactiveMode(l *localizer.Localizer, langs []string) (string, error) {
	prompt := promptui.Select{
		Label: l.Translate("Select the language to use"),
		Items: langs,
	}

	_, result, err := prompt.Run()
	if err != nil {
		slog.Error("Language selection prompt failed", slog.String("error", err.Error()))
		return "", err
	}
	return result, nil
}

func setLanguage(out io.Writer, l *localizer.Localizer, lang string) {
	conf := config.GetConfig()
	fmt.Fprintln(out, l.Translate("You selected: %s", lang))
	err := config.SaveToConfig(conf, "language", lang)
	if err != nil {
		fmt.Fprintln(out, l.Translate("Error saving language to config."))
		os.Exit(1)
	}
}

func setFlags(cmd *cobra.Command, l *localizer.Localizer) {
}

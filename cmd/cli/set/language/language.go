package language

import (
	"fmt"
	"log/slog"

	"github.com/manifoldco/promptui"
	"github.com/mrtnhwttktc/kinto-cli/cmd/utils"
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
		Run: func(cmd *cobra.Command, args []string) {
			langs := localizer.GetLangOptions()
			conf := config.GetConfig()

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
			}
		},
	}
	bindFlags(languageCmd, l)
	return languageCmd
}

func bindFlags(cmd *cobra.Command, l *localizer.Localizer) {
	utils.LocalizeHelpFlag(cmd, l)
}

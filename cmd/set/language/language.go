package language

import (
	"fmt"
	"log/slog"

	"github.com/manifoldco/promptui"
	"github.com/mrtnhwttktc/kinto-cli/internal/config"
	"github.com/mrtnhwttktc/kinto-cli/internal/localizer"
	"github.com/spf13/cobra"
)

func NewLanguageCmd() *cobra.Command {
	l := localizer.GetLocalizer()
	languageCmd := &cobra.Command{
		Use:   "language [english|japanese]",
		Short: "A brief description of your command",
		Long:  `Long description of your command.`,
		Example: `
	# interactive mode
	ktcli config language

	# args mode
	ktcli config language english
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
	return languageCmd
}

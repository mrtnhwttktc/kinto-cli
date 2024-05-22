package cmd

import (
	"fmt"
	"log/slog"

	"github.com/manifoldco/promptui"
	"github.com/mrtnhwttktc/kinto-cli/internal/config"
	"github.com/mrtnhwttktc/kinto-cli/internal/localizer"
	"github.com/spf13/cobra"
)

// languageCmd represents the language command
var languageCmd = &cobra.Command{
	Use:   "language [english|japanese]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Example: `
	# interactive mode
	ktcli config language

	# args mode
	ktcli config language english
	`,
	Run: func(cmd *cobra.Command, args []string) {

		langs := localizer.GetLangOptions()

		prompt := promptui.Select{
			Label: localizer.Lang.Translate("Select the language to use"),
			Items: langs,
		}

		_, result, err := prompt.Run()
		if err != nil {
			slog.Error("Language selection prompt failed", slog.String("error", err.Error()))
			return
		}
		fmt.Println(localizer.Lang.Translate("You selected: %s", result))
		err = config.SaveToConfig("language", result)
		if err != nil {
			slog.Error("Error saving language to config.", slog.String("error", err.Error()))
		}
	},
}

func init() {
	configCmd.AddCommand(languageCmd)
}

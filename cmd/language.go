/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
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
	Use:   "language",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		langs := localizer.GetLangOptions()

		prompt := promptui.Select{
			Label: localizer.Lang.Translate("Select the language to use"),
			Items: langs,
		}

		_, result, err := prompt.Run()
		if err != nil {
			slog.Error("Prompt failed: %v\n", err)
			return
		}
		fmt.Println(localizer.Lang.Translate("You selected: %s", result))
		err = config.SaveToConfig("language", result)
		if err != nil {
			slog.Error("Error saving language to config: %v\n", err)
		}
	},
}

func init() {
	configCmd.AddCommand(languageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// languageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// languageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

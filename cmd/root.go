/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/mrtnhwttktc/kinto-cli/internal/localizer"
	v "github.com/mrtnhwttktc/kinto-cli/internal/version"
	"github.com/spf13/cobra"
)

var version = v.Version

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "ktcli",
	Version: version,
	Short:   localizer.Lang.Translate("CLI for Kinto scripts and tools."),
	Long:    localizer.Lang.Translate(`Kinto CLI or ktcli is a command line interface for employees at Kinto Technologies, allowing easy access to the multiple tools and scripts developped by our teams.`),

	// Uncomment the following line if your bare application
	// has an action associated with it:
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().Bool("debug", false, localizer.Lang.Translate("Sets the log level to debug."))
	rootCmd.PersistentFlags().Bool("verbose", false, localizer.Lang.Translate("Prints logs to stdout."))

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}

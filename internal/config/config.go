package config

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// initDefault sets the default values for the config
func initDefault() {
	viper.SetDefault("language", "japanese")
	viper.SetDefault("log_level", "info")
}

func init() {
	slog.Info("Initializing config")
	viper.GetViper()
	// set the defaults
	initDefault()
	viper.SetEnvPrefix("KTCLI_")
	viper.AutomaticEnv() // read in environment variables that match

	// Find home directory.
	home, err := os.UserHomeDir()
	if err != nil {
		slog.Error("Error getting home directory, please check if the $HOME variable is set. Using the default config.", slog.String("error", err.Error()))
		return
	}

	configDir := filepath.Join(home, ".config", "ktcli")

	viper.AddConfigPath(configDir)
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	// If a config file is found, read it in, else create a new one
	if err := viper.ReadInConfig(); err != nil {
		slog.Info("No config file found, creating a new one.")
		err := writeNewConfigFile()
		if err != nil {
			slog.Error("Error creating a new config file.", slog.String("error", err.Error()))
		}
		return
	}
	slog.Info(fmt.Sprintf("Using config file: %s", viper.ConfigFileUsed()))
}

// writeNewConfigFile creates a new config file in the $HOME/.config/ktcli folder with the default values
func writeNewConfigFile() error {
	home, err := os.UserHomeDir()
	if err != nil {
		slog.Error("Error getting home directory, please check if the $HOME variable is set. Using the default config.", slog.String("error", err.Error()))
		return err
	}

	configDir := filepath.Join(home, ".config", "ktcli")

	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err := os.MkdirAll(home+"/.config/ktcli", 0755)
		if err != nil {
			slog.Error("Error creating the config directory.", slog.String("error", err.Error()))
			return err
		}
	}
	if err := viper.SafeWriteConfigAs(home + "/.config/ktcli/config.yaml"); err != nil {
		return err
	}
	return nil
}

func SaveToConfig(key string, value string) error {
	viper.Set(key, value)
	if err := viper.WriteConfig(); err != nil {
		slog.Error("Error writing to config file.", slog.String("error", err.Error()))
		return err
	}
	return nil
}

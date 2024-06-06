package config

import (
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// initDefault sets the default values for the config
func initializeDefault() {
	viper.SetDefault("language", "japanese")
	viper.SetDefault("log-level", "info")
}

func InitializeConfig() {
	slog.Info("Initializing config")
	// set the defaults
	initializeDefault()

	// load the config file
	if err := loadConfigFile(); err != nil {
		slog.Warn("Failed reading the config file.")
	}

	// set the environment variables
	viper.SetEnvPrefix("KTCLI")
	// Environment variables can't have dashes in them, so bind them to their equivalent
	// keys with underscores, e.g. --favorite-color to KTCLI_FAVORITE_COLOR
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv() // read in environment variables that match

}

// loadConfigFile loads the config file from the default location
// If the file does not exist, a new one is created with the default values
func loadConfigFile() error {
	// Find home directory.
	home, err := os.UserHomeDir()
	if err != nil {
		slog.Error("Error getting home directory, check if the $HOME environment variable is set.", slog.String("error", err.Error()))
		return err
	}

	configDir := filepath.Join(home, ".config", "ktcli")

	viper.AddConfigPath(configDir)
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	// Check if the config file exists
	if _, err := os.Stat(configDir + "/config.yaml"); os.IsNotExist(err) {
		slog.Info("No config file found, creating a new one.")
		err := writeNewConfigFile()
		if err != nil {
			slog.Warn("Error creating a new config file.", slog.String("error", err.Error()))
			return err
		}
		slog.Info("New config file created.", slog.String("path", configDir+"/config.yaml"))
	}

	// If a config file is found, read it in
	if err := viper.ReadInConfig(); err != nil {
		slog.Error("Error reading config file.", slog.String("error", err.Error()))
		return err
	}
	return nil
}

// writeNewConfigFile creates a new config file in the $HOME/.config/ktcli folder with the default values
func writeNewConfigFile() error {
	home, err := os.UserHomeDir()
	if err != nil {
		slog.Warn("Error getting home directory, please check if the $HOME variable is set. Using the default config.", slog.String("error", err.Error()))
		return err
	}

	configDir := filepath.Join(home, ".config", "ktcli")

	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err := os.MkdirAll(home+"/.config/ktcli", 0755)
		if err != nil {
			slog.Warn("Error creating the config directory.", slog.String("error", err.Error()))
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

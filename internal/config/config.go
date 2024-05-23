package config

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var v *viper.Viper

// GetConfig returns the viper config using the singleton pattern
// If the config is not initialized, it will be created and initialized before returning
func GetConfig() *viper.Viper {
	if v == nil {
		v = viper.New()
		initializeConfig(v)
	}
	return v
}

// initDefault sets the default values for the config
func initializeDefault(conf *viper.Viper) {
	conf.SetDefault("language", "japanese")
	conf.SetDefault("log_level", "info")
}

func initializeConfig(conf *viper.Viper) {
	slog.Info("Initializing config")
	// set the defaults
	initializeDefault(conf)
	conf.SetEnvPrefix("KTCLI")
	conf.AutomaticEnv() // read in environment variables that match

	// Find home directory.
	home, err := os.UserHomeDir()
	if err != nil {
		slog.Error("Error getting home directory, please check if the $HOME variable is set. Using the default config.", slog.String("error", err.Error()))
		return
	}

	configDir := filepath.Join(home, ".config", "ktcli")

	conf.AddConfigPath(configDir)
	conf.SetConfigType("yaml")
	conf.SetConfigName("config")

	// If a config file is found, read it in, else create a new one
	if err := conf.ReadInConfig(); err != nil {
		slog.Info("No config file found, creating a new one.")
		err := writeNewConfigFile(conf)
		if err != nil {
			slog.Error("Error creating a new config file.", slog.String("error", err.Error()))
		}
		return
	}
	slog.Info(fmt.Sprintf("Using config file: %s", conf.ConfigFileUsed()))
}

// writeNewConfigFile creates a new config file in the $HOME/.config/ktcli folder with the default values
func writeNewConfigFile(conf *viper.Viper) error {
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
	if err := conf.SafeWriteConfigAs(home + "/.config/ktcli/config.yaml"); err != nil {
		return err
	}
	return nil
}

func SaveToConfig(conf *viper.Viper, key string, value string) error {
	conf.Set(key, value)
	if err := conf.WriteConfig(); err != nil {
		slog.Error("Error writing to config file.", slog.String("error", err.Error()))
		return err
	}
	return nil
}

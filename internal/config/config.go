package config

import (
	"fmt"
	"log/slog"
	"os"

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
	viper.SetEnvPrefix("KTCLI_")
	viper.AutomaticEnv() // read in environment variables that match

	// set the defaults
	initDefault()

	// Find home directory.
	home, err := os.UserHomeDir()
	if err != nil {
		return
	}
	viper.AddConfigPath(home + "/.config/ktcli")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	// If a config file is found, read it in, else create a new one
	if err := viper.ReadInConfig(); err == nil {
		slog.Info(fmt.Sprintf("Using config file: %s", viper.ConfigFileUsed()))
	} else {
		writeNewConfigFile()
	}
}

// writeNewConfigFile creates a new config file in the ~/.config/ktcli folder with the default values
func writeNewConfigFile() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	slog.Warn("No config file found, creating a new one.")
	if _, err := os.Stat(home + "/.config"); os.IsNotExist(err) {
		err := os.Mkdir(home+"/.config", 0755)
		if err != nil {
			return err
		}
	}

	if _, err := os.Stat(home + "/.config/ktcli"); os.IsNotExist(err) {
		err := os.Mkdir(home+"/.config/ktcli", 0755)
		if err != nil {
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
	// check if the config file exists
	// if not, create it with writeNewConfigFile
	// if it fails, return an error
	if err := viper.WriteConfig(); err != nil {
		return err
	}
	return nil
}

package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func TestSaveToConfig(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "ktcli-test")
	require.NoError(t, err, "error creating a temporary test directory")
	testDir, err := os.Getwd()
	require.NoError(t, err, "error getting the current working directory")
	defer os.Chdir(testDir)
	err = os.Chdir(tmpDir)
	require.NoError(t, err, "error changing to the temporary test directory")

	t.Run("Save to config file", func(t *testing.T) {
		err = os.WriteFile(filepath.Join(tmpDir, "test-config.yml"), []byte(""), 0644)
		require.NoError(t, err, "error writing test config file")
		defer os.Remove(filepath.Join(tmpDir, "test-config.yml"))

		viper.New()
		viper.SetConfigFile(filepath.Join(tmpDir, "test-config.yml"))
		viper.SetConfigType("yaml")
		err = viper.ReadInConfig()

		err := SaveToConfig("foo", "bar")
		if err != nil {
			t.Errorf("Error saving to config file: %v", err)
		}
		assert.Equal(t, viper.GetString("foo"), "bar")
		// read the config file again to check if the value was saved
		config, err := os.ReadFile(filepath.Join(tmpDir, "test-config.yml"))
		require.NoError(t, err, "error reading test config file")
		assert.Equal(t, string(config), "foo: bar\n")
	})
}

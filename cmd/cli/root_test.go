package cli

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	v "github.com/mrtnhwttktc/kinto-cli/internal/version"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPrecedence(t *testing.T) {
	testcases := []struct {
		name     string
		language string
		envVar   string
		expected string
	}{
		{
			name:     "Config file set to english, no env var",
			language: "english",
			envVar:   "",
			expected: fmt.Sprintf("ktcli version %s\n", v.Version),
		},
		{
			name:     "Config file set to japanese, no env var",
			language: "japanese",
			envVar:   "",
			expected: fmt.Sprintf("ktcli バージョン %s\n", v.Version),
		},
		{
			name:     "Config file set to english, env var set to japanese",
			language: "english",
			envVar:   "japanese",
			expected: fmt.Sprintf("ktcli バージョン %s\n", v.Version),
		},
		{
			name:     "Config file set to japanese, env var set to english",
			language: "japanese",
			envVar:   "english",
			expected: fmt.Sprintf("ktcli version %s\n", v.Version),
		},
		{
			name:     "Config file set to japanese, env var set to japanese",
			language: "japanese",
			envVar:   "japanese",
			expected: fmt.Sprintf("ktcli バージョン %s\n", v.Version),
		},
	}

	tmpDir, err := os.MkdirTemp("", "ktcli-test")
	require.NoError(t, err, "error creating a temporary test directory")
	testDir, err := os.Getwd()
	require.NoError(t, err, "error getting the current working directory")
	defer os.Chdir(testDir)
	err = os.Chdir(tmpDir)
	require.NoError(t, err, "error changing to the temporary test directory")

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err = os.WriteFile(filepath.Join(tmpDir, "test-config.yml"), []byte(fmt.Sprintf("language: %s", tc.language)), 0644)
			require.NoError(t, err, "error writing test config file")
			defer os.Remove(filepath.Join(tmpDir, "test-config.yml"))

			viper.New()
			viper.SetConfigFile(filepath.Join(tmpDir, "test-config.yml"))
			viper.SetConfigType("yaml")
			err = viper.ReadInConfig()
			require.NoError(t, err, "error reading test config file")

			if tc.envVar != "" {
				viper.Set("language", tc.envVar)
			}
			rootCmd := NewRootCmd()
			output := &bytes.Buffer{}
			rootCmd.SetOut(output)
			rootCmd.SetArgs([]string{"-v"})
			rootCmd.Execute()
			gotOutput := output.String()
			assert.Equal(t, tc.expected, gotOutput, "expected output to match")
		})
	}
}

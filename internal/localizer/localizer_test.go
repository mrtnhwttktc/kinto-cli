package localizer

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestNewLocalizer(t *testing.T) {
	testcases := []struct {
		name     string
		language string
		expected string
	}{
		{
			name:     "Create an english localizer",
			language: "english",
			expected: "english",
		},
		{
			name:     "Create a japanese localizer",
			language: "japanese",
			expected: "japanese",
		},
		{
			name:     "No language set, default to english",
			language: "",
			expected: "english",
		},
	}
	for _, tc := range testcases {
		viper.New()
		viper.Set("language", tc.language)
		localizer := NewLocalizer()
		assert.NotNil(t, localizer)
		assert.Equal(t, tc.expected, localizer.Language)
	}
}

func TestTranslate(t *testing.T) {
	testcases := []struct {
		name     string
		language string
		expected string
	}{
		{
			name:     "Create an english localizer",
			language: "english",
			expected: "Please select a language.",
		},
		{
			name:     "Create a japanese localizer",
			language: "japanese",
			expected: "言語を選択してください。",
		},
		{
			name:     "No language set, default to english",
			language: "",
			expected: "Please select a language.",
		},
	}
	for _, tc := range testcases {
		viper.New()
		viper.Set("language", tc.language)
		localizer := NewLocalizer()
		translated := localizer.Translate("Please select a language.")
		assert.Equal(t, tc.expected, translated)
	}
}

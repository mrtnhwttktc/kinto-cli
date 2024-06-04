package localizer

import (
	"github.com/mrtnhwttktc/kinto-cli/internal/config"
	_ "github.com/mrtnhwttktc/kinto-cli/internal/translations"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Localizer is a struct that holds the printer and the language.
// The Language is the plain text name of the language and the printer is the
// message printer that will be used to translate the strings
type Localizer struct {
	printer  *message.Printer
	Language string
}

var locales = []Localizer{
	{
		Language: "english",
		printer:  message.NewPrinter(language.MustParse("en-GB")),
	},
	{
		Language: "japanese",
		printer:  message.NewPrinter(language.MustParse("ja-JP")),
	},
}

// Lang exposes the initialized Localizer, defaulting to english if no language is set
var localizer *Localizer

// GetLocalizer returns the Localizer using the singleton pattern.
// If the localizer is not set, it will set it to the language set in the configuration or default to english
func GetLocalizer() *Localizer {
	if localizer != nil {
		return localizer
	}
	conf := config.GetConfig()
	configLang := conf.GetString("language")
	for _, locale := range locales {
		if configLang == locale.Language {
			localizer = &locale
			return localizer
		}
	}
	localizer = &Localizer{Language: "english", printer: message.NewPrinter(language.MustParse("en-GB"))}
	return localizer
}

// GetLangOptions returns a list of the available languages in english
func GetLangOptions() []string {
	var langs []string
	for _, locale := range locales {
		langs = append(langs, locale.Language)
	}
	return langs
}

// Translate translates the string to the language set in the Localizer
func (l *Localizer) Translate(key message.Reference, args ...interface{}) string {
	return l.printer.Sprintf(key, args...)
}

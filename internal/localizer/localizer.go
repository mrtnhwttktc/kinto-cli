package localizer

import (
	_ "github.com/mrtnhwttktc/kinto-cli/internal/translations"
	"github.com/spf13/viper"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Localizer is a struct that holds the printer and the ID of the language.
// The ID is the plain text name of the language and the printer is the
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
var Lang Localizer = SetLocalizer()

// SetLocalizer sets the language of the Localizer
func SetLocalizer() Localizer {
	configLang := viper.GetString("language")
	for _, locale := range locales {
		if configLang == locale.Language {
			return locale
		}
	}
	return Localizer{Language: "english", printer: message.NewPrinter(language.MustParse("en-GB"))}
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
func (l Localizer) Translate(key message.Reference, args ...interface{}) string {
	return l.printer.Sprintf(key, args...)
}

// init reads the local configuration and initialize the Localizer
// this init function needs to be called after the viper configuration is loaded in the config package
// so that the language can be set correctly
// func init() {
// 	configLang := viper.GetString("language")
// 	fmt.Println(configLang)
// 	for _, locale := range locales {
// 		if configLang == locale.Language {
// 			Lang = locale
// 		}
// 	}
// 	if Lang == (Localizer{}) {
// 		Lang = Localizer{Language: "english", printer: message.NewPrinter(language.MustParse("en-GB"))}
// 	}

// }

package localizer

import (
	_ "github.com/mrtnhwttktc/kinto-cli/internal/translations"
	"github.com/spf13/viper"

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

// NewLocalizer returns a Localizer. Get the language from viper, defaults to english
func NewLocalizer() *Localizer {
	configLang := viper.GetString("language")
	switch configLang {
	case "english":
		return &Localizer{Language: "english", printer: message.NewPrinter(language.MustParse("en-GB"))}
	case "japanese":
		return &Localizer{Language: "japanese", printer: message.NewPrinter(language.MustParse("ja-JP"))}
	default:
		return &Localizer{Language: "english", printer: message.NewPrinter(language.MustParse("en-GB"))}
	}
}

// GetLangOptions returns a list of the available languages
func (l *Localizer) GetLangOptions() []string {
	return []string{"english", "japanese"}
}

// Translate translates the string to the language set in the Localizer
func (l *Localizer) Translate(key message.Reference, args ...interface{}) string {
	return l.printer.Sprintf(key, args...)
}

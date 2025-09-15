package i18n

import (
	"encoding/json"
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var (
	bundle *i18n.Bundle
)

// Init initializes the i18n system
func Init() error {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	// Load translation files
	translationFiles := []string{
		"locales/active.en.json",
		"locales/active.nl.json",
	}

	for _, file := range translationFiles {
		if _, err := bundle.LoadMessageFile(file); err != nil {
			return fmt.Errorf("failed to load translation file %s: %w", file, err)
		}
	}

	return nil
}

// Localizer creates a new localizer for the given language
func Localizer(lang string) *i18n.Localizer {
	if lang == "" {
		lang = "en"
	}
	return i18n.NewLocalizer(bundle, lang)
}

// Translate translates a message with optional template data
func Translate(localizer *i18n.Localizer, messageID string, templateData map[string]interface{}) (string, error) {
	return localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: templateData,
	})
}

// TranslateString translates a message and returns the string directly (fallback to messageID if error)
func TranslateString(localizer *i18n.Localizer, messageID string, templateData map[string]interface{}) string {
	result, err := Translate(localizer, messageID, templateData)
	if err != nil {
		return messageID
	}
	return result
}
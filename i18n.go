package i18n

import (
	"fmt"
	"io"
)

// TranslateFunc returns the translation of the string identified by translationID.
type TranslateFunc func(translationID string, args ...interface{}) string

// Language describes a language
type Language struct {
	// Code is the language code, e.g. "en" or "de"
	Code string
	// Name is the language's native name, e.g. "English" or "Deutsch"
	Name string
}

// AvailableLanguagesProvider provides a list of available languages.
type AvailableLanguagesProvider interface {
	AvailableLanguages() []Language
}

// messageLoadable can load messages from an io.Reader.
type messageLoadable interface {
	AvailableLanguagesProvider
	// LoadMessages loads translations from a reader. fileName must contain
	// the id of the language.
	LoadMessages(fileName string, r io.Reader) error
}

// Service is the abstract service that provides access to the i18n functions.
type Service interface {
	messageLoadable
	// LoadCategory loads all the messages for the specified category.
	LoadCategory(category string) error
	// Tfunc return the translate function for the first language
	// that has translations.
	Tfunc(language string, additionalLanguages ...string) (TranslateFunc, error)
}

// MessageLoader loads messages into a messageLoadable for all its
// languages.
type MessageLoader interface {
	LoadCategory(category string, ml messageLoadable) error
}

// LoadError is a map of errors that happened while loading messages,
// indexed by language code
type LoadError map[string]error

func (e LoadError) Error() string {
	return fmt.Sprintf("Failed to load some messages: %v", map[string]error(e))
}

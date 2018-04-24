package i18n

import (
	"io"
	"io/ioutil"
	"sync"

	"github.com/nicksnyder/go-i18n/i18n"
	"github.com/pkg/errors"
)

type service struct {
	languages []Language
	mx        sync.RWMutex
	loader    MessageLoader
}

// NewService creates a new i18n service.
func NewService(languages []Language, loader MessageLoader) Service {
	return &service{
		languages: languages,
		loader:    loader,
	}
}

func (s *service) AvailableLanguages() []Language {
	return s.languages
}

func (s *service) LoadMessages(fileName string, r io.Reader) error {
	s.mx.Lock()
	defer s.mx.Unlock()
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return errors.Wrap(err, "failed to read messages:")
	}
	err = i18n.ParseTranslationFileBytes(fileName, b)
	if err != nil {
		return errors.Wrap(err, "failed to parse messages:")
	}
	return nil
}

// Tfunc returns a translator function for the first language which has translations.
func (s *service) Tfunc(language string, additionalLanguages ...string) (TranslateFunc, error) {
	s.mx.RLock()
	defer s.mx.RUnlock()
	t, err := i18n.Tfunc(language, additionalLanguages...)
	return TranslateFunc(t), err
}

func (s *service) LoadCategory(category string) error {
	return s.loader.LoadCategory(category, s)
}

package i18n

import (
	"fmt"
	"os"
	"path/filepath"
)

type fileLoader struct {
	messagePath string
}

// NewFileMessageLoader returns a new MessageLoader that loads
// messages from json files in a directory.
// The files must be named [code].[category].json.
func NewFileMessageLoader(messagePath string) MessageLoader {
	return &fileLoader{
		messagePath: messagePath,
	}
}

func (l *fileLoader) LoadCategory(category string, ml messageLoadable) error {
	var errors LoadError
	for _, language := range ml.AvailableLanguages() {
		fileName := fmt.Sprintf("%v.%v.json", language.Code, category)
		file, err := os.Open(filepath.Join(l.messagePath, fileName))
		if err == nil {
			err = ml.LoadMessages(fileName, file)
		}
		file.Close()
		if err != nil {
			if errors == nil {
				errors = make(LoadError)
			}
			errors[language.Code] = err
		}
	}
	// https://golang.org/doc/faq#nil_error
	if len(errors) > 0 {
		return errors
	}
	return nil
}

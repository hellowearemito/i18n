package i18n

import (
	"bytes"
	"fmt"
)

// ByteSource can provide the contents of named files
// as byte slices.
type ByteSource interface {
	Bytes(name string) ([]byte, error)
}

type bytesLoader struct {
	source ByteSource
}

// NewBytesMessageLoader returns a new MessageLoader that loads
// messages from json files in a ByteSource.
func NewBytesMessageLoader(source ByteSource) MessageLoader {
	return &bytesLoader{
		source: source,
	}
}

func (l *bytesLoader) LoadCategory(category string, ml messageLoadable) error {
	var errors LoadError
	for _, language := range ml.AvailableLanguages() {
		fileName := fmt.Sprintf("%v.%v.json", language.Code, category)
		b, err := l.source.Bytes(fileName)
		buff := bytes.NewBuffer(b)
		if err == nil {
			err = ml.LoadMessages(fileName, buff)
		}
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

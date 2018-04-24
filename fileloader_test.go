package i18n

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

//go:generate counterfeiter -o fakeservice_test.go -pkg ${GOPACKAGE} -fake-name=fakeService . Service

var _ MessageLoader = &fileLoader{}

func TestFileLoader(t *testing.T) {
	Convey("FileMessageLoader", t, func() {
		dir, err := ioutil.TempDir("", "messageloader")
		So(err, ShouldBeNil)
		defer func() {
			os.RemoveAll(dir)
		}()
		err = ioutil.WriteFile(filepath.Join(dir, "en.test.json"), []byte("en messages"), os.ModePerm)
		So(err, ShouldBeNil)
		err = ioutil.WriteFile(filepath.Join(dir, "en-pig.test.json"), []byte("pig latin messages"), os.ModePerm)
		So(err, ShouldBeNil)
		loader := NewFileMessageLoader(dir)
		Convey("Should load files", func() {
			service := new(fakeService)
			service.AvailableLanguagesReturns([]Language{
				{
					Code: "en",
					Name: "english",
				},
				{
					Code: "en-pig",
					Name: "pig latin",
				},
			})
			loadedMessages := make(map[string]string)
			service.LoadMessagesStub = func(fileName string, r io.Reader) error {
				b, err := ioutil.ReadAll(r)
				So(err, ShouldBeNil)
				loadedMessages[fileName] = string(b)
				return nil
			}
			err := loader.LoadCategory("test", service)
			So(err, ShouldBeNil)
			So(service.LoadMessagesCallCount(), ShouldEqual, 2)
			So(loadedMessages["en.test.json"], ShouldEqual, `en messages`)
			So(loadedMessages["en-pig.test.json"], ShouldEqual, `pig latin messages`)
		})

		Convey("Should return load errors", func() {
			service := new(fakeService)
			service.AvailableLanguagesReturns([]Language{
				{
					Code: "en",
					Name: "english",
				},
				{
					Code: "en-pig",
					Name: "pig latin",
				},
			})
			service.LoadMessagesStub = func(fileName string, r io.Reader) error {
				if fileName == "en-pig.test.json" {
					return errors.New("pig latin error")
				}
				return nil
			}
			err := loader.LoadCategory("test", service)
			So(err, ShouldNotBeNil)
			errs, ok := err.(LoadError)
			So(ok, ShouldBeTrue)
			So(errs, ShouldContainKey, "en-pig")
		})

		Convey("Should return file errors", func() {
			service := new(fakeService)
			service.AvailableLanguagesReturns([]Language{
				{
					Code: "en",
					Name: "english",
				},
				{
					Code: "en-pig",
					Name: "pig latin",
				},
				{
					Code: "de",
					Name: "deutsch",
				},
			})
			service.LoadMessagesReturns(nil)
			err := loader.LoadCategory("test", service)
			So(err, ShouldNotBeNil)
			errs, ok := err.(LoadError)
			So(ok, ShouldBeTrue)
			So(errs, ShouldContainKey, "de")
		})
	})
}

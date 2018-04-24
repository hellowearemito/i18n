package i18n

import (
	"errors"
	"io"
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

//go:generate counterfeiter -o fakebytesource_test.go -pkg ${GOPACKAGE} -fake-name=fakeByteSource . ByteSource

var _ MessageLoader = &bytesLoader{}

func TestBytesLoader(t *testing.T) {
	Convey("BytesLoader", t, func() {
		fakeSource := new(fakeByteSource)
		loader := NewBytesMessageLoader(fakeSource)
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
			fakeSource.BytesStub = func(name string) ([]byte, error) {
				return []byte("[" + name + "]"), nil
			}
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
			So(loadedMessages["en.test.json"], ShouldEqual, `[en.test.json]`)
			So(loadedMessages["en-pig.test.json"], ShouldEqual, `[en-pig.test.json]`)
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
			fakeSource.BytesStub = func(name string) ([]byte, error) {
				return []byte("[" + name + "]"), nil
			}
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
	})
}

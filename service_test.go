package i18n

import (
	"bytes"
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

//go:generate counterfeiter -o fakeloader_test.go -pkg ${GOPACKAGE} -fake-name=fakeLoader . MessageLoader

type errorReader struct {
}

func (r *errorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("Test error")
}

func TestService(t *testing.T) {
	Convey("i18n service", t, func() {
		languages := []Language{
			{
				Code: "en",
				Name: "english",
			},
			{
				Code: "en-pig",
				Name: "pig latin",
			},
		}
		loader := new(fakeLoader)
		service := NewService(languages, loader)
		Convey("Should return correct available languages", func() {
			actual := service.AvailableLanguages()
			So(actual, ShouldHaveLength, 2)
			So(actual, ShouldContain, Language{
				Code: "en",
				Name: "english",
			})
			So(actual, ShouldContain, Language{
				Code: "en-pig",
				Name: "pig latin",
			})
		})
		Convey("Should load messages", func() {
			enMessages := bytes.NewBufferString(`[
	{
		"id": "test.message.one",
		"translation": "Test Message One"
	},
	{
		"id": "test.message.two",
		"translation": "Test Message Two"
	}
]
`)
			pigMessages := bytes.NewBufferString(`[
	{
		"id": "test.message.one",
		"translation": "estTay essageMay Oneyay"
	},
	{
		"id": "test.message.two",
		"translation": "estTay essageMay oTway"
	}
]
`)
			err := service.LoadMessages("en.json", enMessages)
			So(err, ShouldBeNil)
			err = service.LoadMessages("en-pig.json", pigMessages)
			So(err, ShouldBeNil)
			Convey("Should translate messages", func() {
				enT, err := service.Tfunc("en")
				So(err, ShouldBeNil)
				pigT, err := service.Tfunc("en-pig")
				So(err, ShouldBeNil)
				So(enT("test.message.one"), ShouldEqual, "Test Message One")
				So(pigT("test.message.one"), ShouldEqual, "estTay essageMay Oneyay")
				So(enT("test.message.two"), ShouldEqual, "Test Message Two")
				So(pigT("test.message.two"), ShouldEqual, "estTay essageMay oTway")
			})
		})

		Convey("Should load category", func() {
			loader.LoadCategoryReturns(nil)
			err := service.LoadCategory("test")
			So(err, ShouldBeNil)
			So(loader.LoadCategoryCallCount(), ShouldEqual, 1)
			category, ml := loader.LoadCategoryArgsForCall(0)
			So(category, ShouldEqual, "test")
			So(ml, ShouldEqual, service)
		})

		Convey("Should fail to parse invalid messages", func() {
			invalid := bytes.NewBufferString(`[{}]`)
			err := service.LoadMessages("en.json", invalid)
			So(err, ShouldNotBeNil)
		})

		Convey("Should fail on reader error", func() {
			err := service.LoadMessages("en.json", &errorReader{})
			So(err, ShouldNotBeNil)
		})
	})
}

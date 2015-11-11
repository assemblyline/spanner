package config_test

import (
	"errors"
	c "github.com/assemblyline/spanner/config"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

type unreadable struct{}

func (u unreadable) Read(p []byte) (i int, err error) {
	return 0, errors.New("could not read, because broken")
}

func TestConfig(t *testing.T) {
	Convey("Loading the spanner config", t, func() {
		fixture, err := os.Open("fixture.toml")
		if err != nil {
			panic(err)
		}
		config, err := c.Read(fixture)
		if err != nil {
			panic(err)
		}

		Convey("Reading Builder Name Metadata", func() {
			So(config.Builder.Name, ShouldEqual, "ruby")
		})

		Convey("Reading Builder Version Metadata", func() {
			So(config.Builder.Version, ShouldEqual, "2.2.3")
		})

		Convey("Reading Step Config", func() {
			So(config.Step[0].Cache, ShouldEqual, "vendor/bundle")
			So(config.Step[0].Script, ShouldResemble, []string{"bundle install -r3 -j4 --path vendor/bundle", "bundle clean"})
			So(config.Step[0].RequiredFiles, ShouldResemble, []string{"Gemfile", "Gemfile.lock"})
		})
	})

	Convey("Errors", t, func() {
		u := unreadable{}
		_, err := c.Read(u)
		So(err.Error(), ShouldEqual, "could not read, because broken")

		Convey("invalid config", func() {
			invalid, _ := os.Open("fixture.broken")
			_, err := c.Read(invalid)
			So(err.Error(), ShouldStartWith, "toml: unmarshal:")
		})
	})

}

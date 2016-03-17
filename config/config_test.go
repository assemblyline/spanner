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

		Convey("Reading Spanner Name Metadata", func() {
			So(config.Spanner.Name, ShouldEqual, "ruby")
		})

		Convey("Reading Spanner Version Metadata", func() {
			So(config.Spanner.Version, ShouldEqual, "2.2.3")
		})

		Convey("Reading Step Config", func() {
			So(config.Step[0].Dir, ShouldEqual, "vendor/bundle")
			So(config.Step[0].Script[0], ShouldResemble, []string{"bundle", "install", "-r3", "-j4", "--path", "vendor/bundle"})
			So(config.Step[0].Script[1], ShouldResemble, []string{"bundle", "clean"})
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

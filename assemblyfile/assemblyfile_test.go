package assemblyfile_test

import (
	"errors"
	af "github.com/assemblyline/spanner/assemblyfile"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

type unreadable struct{}

func (u unreadable) Read(p []byte) (i int, err error) {
	return 0, errors.New("could not read, because broken")
}

func TestAssemblyfile(t *testing.T) {
	Convey("Loading an Assemblyfile from disk is correct", t, func() {
		assemblyfile, _ := os.Open("Assemblyfile.fixture")
		config, _ := af.Read(assemblyfile)

		Convey("The application has a name and a repo", func() {
			So(config.Application.Name, ShouldEqual, "Test App")
			So(config.Application.Repo, ShouldEqual, "foo.example.com/assemblyline/test")
		})

		Convey("The build has a builder and a version", func() {
			So(config.Build.Builder, ShouldEqual, "ruby")
			So(config.Build.Version, ShouldEqual, "2.2.3")
		})

		Convey("The test has a script and an ENV", func() {
			So(config.Test.Script, ShouldResemble, []string{"bundle exec rake db:test:prepare", "bundle exec rake"})
			So(config.Test.Env, ShouldResemble, map[string]interface{}{"RACK_ENV": "test", "AWESOME": true})
		})

		Convey("The test has Services", func() {
			So(len(config.Test.Service), ShouldEqual, 3)
			So(config.Test.Service["postgres"].Version, ShouldEqual, "9.4.1")
			So(config.Test.Service["elasticsearch"].Properties[0], ShouldEqual, "es.script.groovy.sandbox.enabled=true")
		})

		Convey("Errors", func() {
			Convey("io error", func() {
				u := unreadable{}
				_, err := af.Read(u)
				So(err.Error(), ShouldEqual, "could not read, because broken")
			})

			Convey("invalid assemblyfile", func() {
				assemblyfile, _ := os.Open("Assemblyfile.broken")
				_, err := af.Read(assemblyfile)
				So(err.Error(), ShouldStartWith, "toml: unmarshal:")
			})
		})
	})

}

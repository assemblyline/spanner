package assemblyfile

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAssemblyfile(t *testing.T) {
	Convey("Loading an Assemblyfile from disk is correct", t, func() {
		assemblyfile := Read("Assemblyfile.fixture")

		Convey("The application has a name and a repo", func() {
			So(assemblyfile.Application.Name, ShouldEqual, "Test App")
			So(assemblyfile.Application.Repo, ShouldEqual, "foo.example.com/assemblyline/test")
		})

		Convey("The build has a builder and a version", func() {
			So(assemblyfile.Build.Builder, ShouldEqual, "ruby")
			So(assemblyfile.Build.Version, ShouldEqual, "2.2.3")
		})

		Convey("The test has a script and an ENV", func() {
			So(assemblyfile.Test.Script, ShouldResemble, []string{"bundle exec rake db:test:prepare", "bundle exec rake"})
			So(assemblyfile.Test.Env, ShouldResemble, map[string]interface{}{"RACK_ENV": "test", "AWESOME": true})
		})

		Convey("The test has Services", func() {
			So(len(assemblyfile.Test.Service), ShouldEqual, 3)
			So(assemblyfile.Test.Service["postgres"].Version, ShouldEqual, "9.4.1")
			So(assemblyfile.Test.Service["elasticsearch"].Properties[0], ShouldEqual, "es.script.groovy.sandbox.enabled=true")
		})
	})
}

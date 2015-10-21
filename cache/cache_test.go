package cache

import (
	"github.com/assemblyline/spanner/assemblyfile"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func config(appName string, appRepo string, builder string, builderVersion string) assemblyfile.Config {
	return assemblyfile.Config{
		Application: assemblyfile.Application{
			Name: appName,
			Repo: appRepo,
		},
		Build: assemblyfile.Build{
			Builder: builder,
			Version: builderVersion,
		},
	}
}

func TestCache(t *testing.T) {
	Convey("Cache key is based on the Assemblyfile and Dir", t, func() {
		baseConfig := config("awesome app", "dockerhub.foo/bar/whathahver", "ruby", "2.3.3")
		baseHash := New("/foo", baseConfig).Hash

		So(baseHash, ShouldEqual, "5d773b69617d6214612fe76c2065f91124b9d43367b7fafc3984640c70dd4a0c")

		Convey("The key is changed if the dir changes", func() {
			So(New("/bar", baseConfig), ShouldNotEqual, baseHash)
		})

		Convey("The key is changed if the application name changes", func() {
			changedConfig := baseConfig
			changedConfig.Application.Name = "super app"
			So(New("/foo", baseConfig), ShouldNotEqual, baseHash)
		})
	})
}

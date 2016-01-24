package cache

import (
	"github.com/assemblyline/spanner/assemblyfile"
	"github.com/assemblyline/spanner/logger"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"os"
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
	Convey("With the FileStore Cache", t, func() {

		cacheDir, _ := ioutil.TempDir("", "cache_test")
		testDir, _ := ioutil.TempDir("", "test_dir")
		data := []byte("hello cache")
		ioutil.WriteFile(testDir+"/foo.txt", data, 0644)
		ioutil.WriteFile(testDir+"/bar.txt", data, 0755)

		cfg := config(
			"Test App",
			"foo.example.com/assemblyline/test",
			"ruby",
			"2.2.3",
		)

		fs := NewFileStore(cacheDir)
		c := New(cfg.Hash(), fs)
		c.log = logger.TestLogger()
		task := "test"

		Convey("Save and Restore with the FileStore Cache", func() {
			c.Save(testDir, task)

			// Simulate a fresh build with a clean checkout
			os.RemoveAll(testDir)
			c.Restore(testDir, task)

			info, err := os.Stat(testDir + "/foo.txt")
			So(err, ShouldBeNil)
			So(info.Mode(), ShouldEqual, 0644)

			readData, _ := ioutil.ReadFile(testDir + "/foo.txt")
			So(string(readData), ShouldEqual, "hello cache")

			info, err = os.Stat(testDir + "/bar.txt")
			So(err, ShouldBeNil)
			So(info.Mode(), ShouldEqual, 0755)

			readData, _ = ioutil.ReadFile(testDir + "/bar.txt")
			So(string(readData), ShouldEqual, "hello cache")
		})

		Convey("Reading from the FileStore Cache when a save has not been made", func() {
			c.Restore(testDir, task)

			// It does not disturb whatever happens to be in the dir
			info, err := os.Stat(testDir + "/foo.txt")
			So(err, ShouldBeNil)
			So(info.Mode(), ShouldEqual, 0644)

			readData, _ := ioutil.ReadFile(testDir + "/foo.txt")
			So(string(readData), ShouldEqual, "hello cache")

			info, err = os.Stat(testDir + "/bar.txt")
			So(err, ShouldBeNil)
			So(info.Mode(), ShouldEqual, 0755)

			readData, _ = ioutil.ReadFile(testDir + "/bar.txt")
			So(string(readData), ShouldEqual, "hello cache")
		})

		Convey("Save errors will panic", func() {
			fs := FileStore{dir: "/does/not/exist"}
			c := New(cfg.Hash(), fs)
			So(func() { c.Save(testDir, task) }, ShouldPanic)
		})

		Reset(func() {
			os.RemoveAll(testDir)
			os.RemoveAll(cacheDir)
		})
	})
}

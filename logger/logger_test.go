package logger

import (
	"bytes"
	"github.com/mgutz/ansi"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestCache(t *testing.T) {
	Convey("logging", t, func() {
		out := &bytes.Buffer{}
		err := &bytes.Buffer{}
		log := Logger{
			out: out,
			err: err,
		}

		Convey("New sets up logger with standard out and err", func() {
			newLogger := New()

			So(newLogger.out, ShouldEqual, os.Stdout)
			So(newLogger.err, ShouldEqual, os.Stderr)
		})

		Convey("titles", func() {
			log.Title("foo", "bar", "baz")
			So(out.String(), ShouldContainSubstring, ansi.Color("[ foo bar baz ]", "black+b:yellow"))
		})

		Convey("info", func() {
			log.Info("foo", "bar", "baz")
			So(out.String(), ShouldContainSubstring, ansi.Color("foo bar baz", "blue"))
		})

		Convey("error", func() {
			log.Error("foo", "bar", "baz")
			So(err.String(), ShouldContainSubstring, ansi.Color("foo bar baz", "red"))
		})
	})
}

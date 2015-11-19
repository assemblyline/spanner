package logger

import (
	"bytes"
	"github.com/mgutz/ansi"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestLogging(t *testing.T) {
	Convey("logging", t, func() {
		log := TestLogger()

		Convey("New sets up logger with standard out and err", func() {
			newLogger := New()

			So(newLogger.out, ShouldEqual, os.Stdout)
			So(newLogger.err, ShouldEqual, os.Stderr)
			So(newLogger.Out(), ShouldEqual, os.Stdout)
			So(newLogger.Err(), ShouldEqual, os.Stderr)
		})

		Convey("titles", func() {
			log.Title("foo", "bar", "baz")
			So(log.out.(*bytes.Buffer).String(), ShouldContainSubstring, ansi.Color("[ foo bar baz ]", "black+b:yellow"))
		})

		Convey("step titles", func() {
			log.StepTitle("foo", "bar", "baz")
			So(log.out.(*bytes.Buffer).String(), ShouldContainSubstring, ansi.Color("==>   foo bar baz   ", "black+b:cyan"))
		})

		Convey("info", func() {
			log.Info("foo", "bar", "baz")
			So(log.out.(*bytes.Buffer).String(), ShouldContainSubstring, ansi.Color("foo bar baz", "blue"))
		})

		Convey("error", func() {

			log.Error("foo", "bar", "baz")
			So(log.err.(*bytes.Buffer).String(), ShouldContainSubstring, ansi.Color("foo bar baz", "red"))
		})
	})
}

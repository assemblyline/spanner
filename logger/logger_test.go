package logger_test

import (
	"bytes"
	"github.com/assemblyline/spanner/logger"
	"github.com/mgutz/ansi"
	"io"
	"os"
	"testing"
)

func TestNewLogger(t *testing.T) {
	newLogger := logger.New()

	if newLogger.Out() != os.Stdout {
		t.Error("Expected logger output to be stdout")
	}

	if newLogger.Err() != os.Stderr {
		t.Error("Expected logger err to be stderr")
	}
}

func TestLoggingTitles(t *testing.T) {
	log := logger.Test()
	log.Title("foo", "bar", "baz")
	assetLogged(t, log.Out(), ansi.Color("[ foo bar baz ]", "black+b:yellow")+"\n")
}

func TestLoggingStepTitles(t *testing.T) {
	log := logger.Test()
	log.StepTitle("foo", "bar", "baz")
	assetLogged(t, log.Out(), "\n"+ansi.Color("==>   foo bar baz   ", "black+b:cyan")+"\n")
}

func TestLoggingInfo(t *testing.T) {
	log := logger.Test()
	log.Info("foo", "bar", "baz")
	assetLogged(t, log.Out(), ansi.Color("foo bar baz", "blue")+"\n")
}

func TestLoggingError(t *testing.T) {
	log := logger.Test()
	log.Error("foo", "bar", "baz")
	assetLogged(t, log.Err(), ansi.Color("foo bar baz", "red")+"\n")
}

func assetLogged(t *testing.T, l io.Writer, expected string) {
	actual := l.(*bytes.Buffer).String()
	if actual != expected {
		t.Error("Expected", actual, "to be", expected)
	}
}

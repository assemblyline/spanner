package logger_test

import (
	"bytes"
	"github.com/assemblyline/spanner/logger"
	"github.com/mgutz/ansi"
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
	actual := log.Out().(*bytes.Buffer).String()
	expected := ansi.Color("[ foo bar baz ]", "black+b:yellow") + "\n"
	if actual != expected {
		t.Error("Expected", actual, "to be", expected)
	}
}

func TestLoggingStepTitles(t *testing.T) {
	log := logger.Test()
	log.StepTitle("foo", "bar", "baz")
	actual := log.Out().(*bytes.Buffer).String()
	expected := "\n" + ansi.Color("==>   foo bar baz   ", "black+b:cyan") + "\n"
	if actual != expected {
		t.Error("Expected", actual, "to be", expected)
	}
}

func TestLoggingInfo(t *testing.T) {
	log := logger.Test()
	log.Info("foo", "bar", "baz")
	actual := log.Out().(*bytes.Buffer).String()
	expected := ansi.Color("foo bar baz", "blue") + "\n"
	if actual != expected {
		t.Error("Expected", actual, "to be", expected)
	}
}

func TestLoggingError(t *testing.T) {
	log := logger.Test()
	log.Error("foo", "bar", "baz")
	actual := log.Err().(*bytes.Buffer).String()
	expected := ansi.Color("foo bar baz", "red") + "\n"
	if actual != expected {
		t.Error("Expected", actual, "to be", expected)
	}
}

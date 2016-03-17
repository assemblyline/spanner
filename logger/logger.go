package logger

import (
	"bytes"
	"fmt"
	"github.com/mgutz/ansi"
	"io"
	"os"
	"strings"
)

var title = ansi.ColorFunc("black+b:yellow")
var stepTitle = ansi.ColorFunc("black+b:cyan")
var info = ansi.ColorFunc("blue")
var errors = ansi.ColorFunc("red")

// Logger reresents a device to write log messages to
type Logger struct {
	out io.Writer
	err io.Writer
}

// New returns a new instance of Logger that logs to stdout and stderr
func New() Logger {
	return Logger{
		out: os.Stdout,
		err: os.Stderr,
	}
}

// Test returns a new instance of Logger that logs to in memory buffers, for testing.
func Test() Logger {
	return Logger{
		out: &bytes.Buffer{},
		err: &bytes.Buffer{},
	}
}

// Out returns the stdout interface of the logger
func (l Logger) Out() io.Writer {
	return l.out
}

// Err returns the stderr interface of the logger
func (l Logger) Err() io.Writer {
	return l.err
}

// Title logs title test to stdout with appropirate formatting
func (l Logger) Title(text ...string) {
	text = append([]string{"["}, text...)
	text = append(text, "]")
	fmt.Fprintln(l.out, title(strings.Join(text, " ")))
}

// Info logs informational test to stdout with appropirate formatting
func (l Logger) Info(text ...string) {
	fmt.Fprintln(l.out, info(strings.Join(text, " ")))
}

// Error logs informational test to sterr with appropirate formatting
func (l Logger) Error(text ...string) {
	fmt.Fprintln(l.err, errors(strings.Join(text, " ")))
}

// StepTitle logs title test to stdout with appropirate formatting for build step titles
func (l Logger) StepTitle(text ...string) {
	text = append([]string{"==>  "}, text...)
	text = append(text, "  ")
	fmt.Fprintln(l.out, "\n"+stepTitle(strings.Join(text, " ")))
}

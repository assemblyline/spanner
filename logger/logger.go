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

type Logger struct {
	out io.Writer
	err io.Writer
}

func New() Logger {
	return Logger{
		out: os.Stdout,
		err: os.Stderr,
	}
}

func TestLogger() Logger {
	return Logger{
		out: &bytes.Buffer{},
		err: &bytes.Buffer{},
	}
}

func (l Logger) Out() io.Writer {
	return l.out
}

func (l Logger) Err() io.Writer {
	return l.err
}

func (l Logger) Title(text ...string) {
	text = append([]string{"["}, text...)
	text = append(text, "]")
	fmt.Fprintln(l.out, title(strings.Join(text, " ")))
}

func (l Logger) Info(text ...string) {
	fmt.Fprintln(l.out, info(strings.Join(text, " ")))
}

func (l Logger) Error(text ...string) {
	fmt.Fprintln(l.err, errors(strings.Join(text, " ")))
}

func (l Logger) StepTitle(text ...string) {
	text = append([]string{"==>  "}, text...)
	text = append(text, "  ")
	fmt.Fprintln(l.out, "\n"+stepTitle(strings.Join(text, " ")))
}

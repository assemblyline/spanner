package logger

import (
	"fmt"
	"github.com/mgutz/ansi"
	"io"
	"os"
	"strings"
)

var title = ansi.ColorFunc("black+b:yellow")
var info = ansi.ColorFunc("blue")

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

func (l Logger) Title(text ...string) {
	text = append([]string{"["}, text...)
	text = append(text, "]")
	fmt.Fprintln(l.out, title(strings.Join(text, " ")))
}

func (l Logger) Info(text ...string) {
	fmt.Fprintln(l.out, info(strings.Join(text, " ")))
}

package logger

import (
	"fmt"
	"github.com/mgutz/ansi"
	"strings"
)

var title = ansi.ColorFunc("black+b:yellow")
var info = ansi.ColorFunc("blue")

func LogTitle(text ...string) {
	text = append([]string{"["}, text...)
	text = append(text, "]")
	fmt.Println(title(strings.Join(text, " ")))
}

func LogInfo(text ...string) {
  fmt.Println(info(strings.Join(text, " ")))
}

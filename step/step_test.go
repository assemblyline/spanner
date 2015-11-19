package step

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStep(t *testing.T) {
	Convey("running steps", t, func() {
		step := Step{
			Script: [][]string{[]string{"echo", "hello"}},
		}

		step.Exec()
	})
}

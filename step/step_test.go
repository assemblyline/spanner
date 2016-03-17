package step_test

import (
	"github.com/assemblyline/spanner/step"
	"testing"
	//"io/ioutil"
)

func TestSimpleStepExec(t *testing.T) {
	s := step.Step{Script: [][]string{[]string{"echo", "hello"}}}
	err := s.Exec()
	if err != nil {
		t.Error(err.Error())
	}
}

package silog

import (
	"errors"
	"testing"
)

func TestLogz(t *testing.T) {
	l := NewLogz("default")
	m := make(map[string]string)
	m["1"] = "hah"
	m["2"] = "haeh"
	defer l.Exit(0, "why need exit in test ?")
	l.Debug("haha")
	l.Debug("[haha]")
	l.Info(m)
	l.Warn([]int{1, 2, 3})
	l.Error(errors.New("eroor"))
}

package logz

import (
	"errors"
	"testing"
	//	"fmt"
)

type testlog struct {
	A int
	B int
	C int
}

func TestPrintln(t *testing.T) {

	ll := NewLogDebug(true)
	ll.Error(errors.New("test err"), "message")
	ll.Info("request", "hello")
	ll.Println("1", "2")
	a := &testlog{1, 2, 3}
	ll.Println(*a)
	b := map[int]int{
		1: 2,
		2: 3,
		4: 4,
	}
	ll.Println(b)
}

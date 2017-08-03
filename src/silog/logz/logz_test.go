package logz

import "testing"

func TestLogz(t *testing.T) {
	l := NewLogz("default")
	//TODO:问题：为什么吧defer注释掉，就没有信息打印？
	defer l.Exit(0, "exit")
	l.Debug("haha")
}

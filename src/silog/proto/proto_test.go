package proto

import (
	"fmt"
	"testing"
)

func TestProto(t *testing.T) {
	protocal := &DefaultProto{
		Level:   Level(1),
		LogName: "default",
	}

	protocal.ToMsg(1, "31231", "haha")

	fmt.Println(string(protocal.Msg))
}

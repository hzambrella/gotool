package netTool

import (
	"fmt"
	"testing"
)

func TestIP(t *testing.T) {
	fmt.Println(GetIp())
	ns, err := GetIpByUrl("www.baidu.com")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ns)
}

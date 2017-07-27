package cloudprt

import (
	"fmt"
	"testing"
)

func TestPrintContent(t *testing.T) {
	prt, err := NewPrt(S1, "kdt1080249", "c1600")
	if err != nil {
		t.Fatal(err)
	}
	_, err = prt.GetPrtStatus()
	if err != nil {
		t.Fatal(err)
	} else {
		resp, err := prt.PrintString("haha")
		if err != nil {
			t.Fatal(err)
		} else {
			qresp, err := prt.QueryOrder(resp.OrderIndex)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Println("order:", qresp)
		}
	}

}

var prtCnt string = `
我们假设你已经有一种或多种其他编程语言的使用经历，不管是类似C、c++或Java的编译型语言，还是类似Python、Ruby、JavaScript的脚本语言，因此我们不会像对完全的编程语言初学者那样解释所有的细节。因为Go语言的变量、常量、表达式、控制流和函数等基本语法也是类似的。
必要时，Go语言工具会创建目录。例如：
`

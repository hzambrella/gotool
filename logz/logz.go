package logz

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

const (
	TIME_FORMAT = "2006-01-02 15:04"
)

type LogDebug struct {
	IfPrintln bool
}

//var IfPrintln bool=false

// 代码复制到项目即可，不需要导入包。
//TODO：log serve
// logDebug or not,true is on ,false is off
//var ifprint bool = false
//var logz = NewLogDebug(ifprint)

func NewLogDebug(ifprintln bool) *LogDebug {
	return &LogDebug{ifprintln}
}

// about depth:please find api :package runtime
//output func name ,file name,line num
func fileLine(depth int) (string, string, int) {
	pc, file, line, ok := runtime.Caller(depth)
	var funcname string = ""
	if !ok {
		file = "?"
		line = 0
	} else {
		funcname = runtime.FuncForPC(pc).Name()
	}
	return funcname, file, line
}

// 断点调试，显示行号和列号,以及函数名，方便断点调试
//结构体传指针，能打出成员名字。
func (ll *LogDebug) Println(t ...interface{}) {
	/*
		if !IfPrintln{
			return
		}
	*/

	funcname, file, line := fileLine(2)

	slash := strings.LastIndex(file, "/")
	if slash >= 0 {
		file = file[slash+1:]
	}

	slash2 := strings.LastIndex(funcname, "/")
	if slash2 >= 0 {
		funcname = funcname[slash2+1:]
	}

	fmt.Printf("%c[1;42;37m%s%c[0m", 0x1B, "[PRINT]", 0x1B)
	//fmt.Printf("%s:%d:'%s()'",file,line,funcname)
	fmt.Printf(" %c[1;40;36m%s:%c[0m", 0x1B, file, 0x1B)
	fmt.Printf(" %c[1;40;35m%d%c[0m", 0x1B, line, 0x1B)
	fmt.Printf(" %c[1;40;34m%s()%c[0m", 0x1B, funcname, 0x1B)
	//	fmt.Printf("[[PRINT] 0x1B[1;40;32m%s0x1B[0m :%d:'%s()']",file,line,funcname)
	fmt.Printf("--->")
	fmt.Println(fmt.Sprintf("%+v\n", t))

}

func (ll *LogDebug) Info(info string, t ...interface{}) {
	now := time.Now()

	/*
		funcname, file, line := fileLine(2)

		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}

		slash2 := strings.LastIndex(funcname, "/")
		if slash2 >= 0 {
			funcname = funcname[slash2+1:]
		}
	*/

	fmt.Printf("%c[1;44;37m%s[%s]%c[0m", 0x1B, now.Format(TIME_FORMAT), "INFO", 0x1B)
	fmt.Printf(" %c[1;40;36m%s:%c[0m", 0x1B, info, 0x1B)
	//fmt.Printf(" %c[1;40;35m%d%c[0m", 0x1B, line, 0x1B)
	//fmt.Printf(" %c[1;40;34m%s()%c[0m", 0x1B, funcname, 0x1B)
	//	fmt.Printf("--->")
	fmt.Println(t)
	//	fmt.Printf("[%s,ERROR]:%s:%d\n",now.Format(TIME_FORMAT),file,line)
}

// 错误输出。带行号和列号,以及函数名，方便检查错误和异常
func (ll *LogDebug) Error(err error, t ...interface{}) {
	if err == nil {
		return
	}
	now := time.Now()

	funcname, file, line := fileLine(2)

	slash := strings.LastIndex(file, "/")
	if slash >= 0 {
		file = file[slash+1:]
	}

	slash2 := strings.LastIndex(funcname, "/")
	if slash2 >= 0 {
		funcname = funcname[slash2+1:]
	}

	fmt.Printf("%c[1;41;37m%s%s%c[0m", 0x1B, now.Format(TIME_FORMAT), "[ERROR]", 0x1B)
	fmt.Printf(" %c[1;40;36m%s:%c[0m", 0x1B, file, 0x1B)
	fmt.Printf(" %c[1;40;35m%d%c[0m", 0x1B, line, 0x1B)
	fmt.Printf(" %c[1;40;34m%s()%c[0m", 0x1B, funcname, 0x1B)
	fmt.Printf("--->")
	fmt.Printf(" %c[1;40;32m%v%c[0m\n", 0x1B, err, 0x1B)
	fmt.Println(fmt.Sprintf("%+v\n", t))
	//	fmt.Printf("[%s,ERROR]:%s:%d\n",now.Format(TIME_FORMAT),file,line)
}

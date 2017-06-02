package lognb

import(
	"fmt"
	"runtime"
	"time"
	"strings"
)

const(
	TIME_FORMAT="2006-01-02 15:04"
)

//var IfPrintln bool=false

// about depth:please find api :package runtime
//output func name ,file name,line num
func fileLine(depth int)(string,string,int){
	pc,file,line,ok:=runtime.Caller(depth)
	var funcname string=""
	if !ok{
		file="?"
		line=0
	}else{
		funcname=runtime.FuncForPC(pc).Name()
	}
	return funcname,file,line
}

// println，显示行号和列号，方便调试
func Println(t ...interface{}){
	/*
	if !IfPrintln{
		return
	}
	*/

	funcname,file,line:=fileLine(2)

	slash:=strings.LastIndex(file,"/")
	if slash>=0{
		file=file[slash+1:]
	}

	slash2:=strings.LastIndex(funcname,"/")
	if slash2>=0{
		funcname=funcname[slash2+1:]
	}

	fmt.Printf("%c[1;42;37m%s%c[0m", 0x1B, "[PRINT]", 0x1B)
	//fmt.Printf("%s:%d:'%s()'",file,line,funcname)
	fmt.Printf(" %c[1;40;36m%s:%c[0m", 0x1B, file, 0x1B)
	fmt.Printf(" %c[1;40;35m%d%c[0m", 0x1B, line, 0x1B)
	fmt.Printf(" %c[1;40;34m%s()%c[0m", 0x1B, funcname, 0x1B)
//	fmt.Printf("[[PRINT] 0x1B[1;40;32m%s0x1B[0m :%d:'%s()']",file,line,funcname)
	fmt.Printf("--->")
	fmt.Printf(" %c[1;40;32m%s%c[0m\n", 0x1B, t, 0x1B)

}
// 错误输出。带行号，方便检查错误和异常
func Error(err error){
	if err==nil{
		return
	}
	now:=time.Now()

	funcname,file,line:=fileLine(2)

	slash:=strings.LastIndex(file,"/")
	if slash>=0{
		file=file[slash+1:]
	}

	slash2:=strings.LastIndex(funcname,"/")
	if slash2>=0{
		funcname=funcname[slash2+1:]
	}

	fmt.Printf("%c[1;41;37m%s%s%c[0m", 0x1B,now.Format(TIME_FORMAT),"[ERROR]",0x1B)
	fmt.Printf(" %c[1;40;36m%s:%c[0m", 0x1B, file, 0x1B)
	fmt.Printf(" %c[1;40;35m%d%c[0m", 0x1B, line, 0x1B)
	fmt.Printf(" %c[1;40;34m%s()%c[0m", 0x1B, funcname, 0x1B)
	fmt.Printf("--->")
	fmt.Printf(" %c[1;40;32m%v%c[0m\n", 0x1B, err, 0x1B)
//	fmt.Printf("[%s,ERROR]:%s:%d\n",now.Format(TIME_FORMAT),file,line)
}

package main

import (
	"fmt"
	"runtime"
)

func main() {
	myFunc()
}

func myFunc() {
	pc, _, _, ok := runtime.Caller(1)
	if ok {
		f := runtime.FuncForPC(pc)
		fmt.Println("调用的函数名称：", f.Name())
	}
}

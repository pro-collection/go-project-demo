package main

import (
	"go-project-demo/packages/pro3_proxypool/pkg/initial"
	"unknwon.dev/clog/v2"
)

func deferExec() {
	defer clog.Stop()
}

func main() {
	// 初始化
	initial.GlobalInit()

	deferExec()
}

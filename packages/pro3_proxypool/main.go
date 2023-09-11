package main

import (
	"go-project-demo/packages/pro3_proxypool/api"
	"go-project-demo/packages/pro3_proxypool/pkg/initial"
	"runtime"
	"unknwon.dev/clog/v2"
)

func deferExec() {
	defer clog.Stop()
}

func main() {
	// 初始化
	initial.GlobalInit()

	runtime.GOMAXPROCS(runtime.NumCPU())

	//ipChan := make(chan *models.IP, 2000)

	go func() {
		api.Run()
	}()

	deferExec()
}

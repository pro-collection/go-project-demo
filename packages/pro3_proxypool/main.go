package main

import (
	"go-project-demo/packages/pro3_proxypool/api"
	"go-project-demo/packages/pro3_proxypool/pkg/initial"
	"go-project-demo/packages/pro3_proxypool/pkg/utils/handleFile"
	"unknwon.dev/clog/v2"
)

func deferExec() {
	defer clog.Stop()
}

func main() {
	// 初始化
	initial.GlobalInit()

	handleFile.GetIP()

	api.RunWithLocal()

	deferExec()
}

package main

import (
	"fmt"
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

	file, err := handleFile.FindFile("source_ip.json")
	defer file.Close()
	if err != nil {
		return
	}

	// 网络获取 ip 地址核心方法
	handleFile.WriteFileWithNetWork(file)

	// 重新读取一遍文档
	file, err = handleFile.FindFile("source_ip.json")
	fileContent, err := handleFile.ReadFile(file)

	list := handleFile.FilterGetUsedIpList(fileContent)
	fmt.Println("ip 处理结束")
	for _, ip := range list {
		fmt.Println("ip: ", ip)
	}

	//重新写入文件
	handleFile.WriteToLocal("ip.json", &list)
}

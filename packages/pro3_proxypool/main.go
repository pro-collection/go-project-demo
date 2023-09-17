package main

import (
	"fmt"
	"go-project-demo/packages/pro3_proxypool/pkg/utils/handleFile"
	"unknwon.dev/clog/v2"
)

func deferExec() {
	defer clog.Stop()
}

func main() {
	// 初始化
	//initial.GlobalInit()

	//runtime.GOMAXPROCS(runtime.NumCPU())

	//ipChan := make(chan *models.IP, 2000)
	//
	//api.Run()
	////go func() {
	////	api.Run()
	////}()
	//
	//storage.CheckProxyDB()
	//
	//for i := 0; i < 50; i++ {
	//	go func() {
	//		for {
	//			storage.CheckProxy(<-ipChan)
	//		}
	//	}()
	//}
	//
	//for {
	//	n := models.CountIps()
	//	logger.Info(&logger.Params{
	//		Key:      logger.Key.BaseInfo,
	//		ModeName: "main",
	//		FuncName: "main",
	//		Content:  fmt.Sprintf("Chan: %v, IP: %v\n", len(ipChan), n),
	//	})
	//
	//	if len(ipChan) < 100 {
	//		// todo yanlele
	//	}
	//
	//	time.Sleep(10 * time.Minute)
	//}
	//
	//deferExec()

	file, err := handleFile.FindFile("ip.json")
	defer file.Close()
	if err != nil {
		return
	}

	//handleFile.WriteFile(file)

	fileContent, err := handleFile.ReadFile(file)

	list := handleFile.FilterGetUsedIpList(fileContent)
	fmt.Println("ip 处理结束")
	for _, ip := range list {
		fmt.Println("ip: ", ip)
	}
}

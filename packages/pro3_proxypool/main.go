package main

import (
	"encoding/json"
	"fmt"
	"go-project-demo/packages/pro3_proxypool/pkg/models"
	"go-project-demo/packages/pro3_proxypool/pkg/storage"
	"go-project-demo/packages/pro3_proxypool/pkg/utils/handleFile"
	"sync"
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

	fileInfo, err := handleFile.ReadFile(file)

	var ipList []*models.IP
	err = json.Unmarshal(fileInfo, &ipList)
	if err != nil {
		fmt.Println("error: ", err.Error())
	}

	var wg sync.WaitGroup

	// 校验 ip 是否好用
	for _, ip := range ipList {
		fmt.Println("ip: ", ip.Data)
		wg.Add(1)
		go func(value *models.IP) {
			used := storage.CheckIP(value)
			fmt.Printf("ip: %s, is used: %t\n", value.Data, used)
			wg.Done()
		}(ip)
	}

	wg.Wait()
}

// todo yanlele run
//func run(ipChan chan<- *models.IP) {
//	var wg sync.WaitGroup
//
//	for _, f := range funs {
//		wg.Add(1)
//		go func(f func() []*models.IP) {
//			temp := f()
//			//log.Println("[run] get into loop")
//			for _, v := range temp {
//				//log.Println("[run] len of ipChan %v",v)
//				ipChan <- v
//			}
//			wg.Done()
//		}(f)
//	}
//	wg.Wait()
//	log.Println("All getters finished.")
//}

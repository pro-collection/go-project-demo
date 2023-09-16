package main

import (
	"go-project-demo/packages/pro3_proxypool/pkg/utils"
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

	//res := getter.IP89()
	//for _, ip := range res {
	//	fmt.Printf("ip: %s\n", ip.Data)
	//}

	// 获取当前工作目录的绝对路径
	//wd, err := os.Getwd()
	//if err != nil {
	//	fmt.Println("获取当前工作目录失败:", err)
	//	return
	//}
	//
	//// 获取相对路径
	//relativePath := "subdir/example.txt"
	//absPath := filepath.Join(wd, relativePath)
	//
	//fmt.Println("相对路径:", relativePath)
	//fmt.Println("绝对路径:", absPath)

	_, err := utils.FindFile("ip.json")
	if err != nil {
		return
	}
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

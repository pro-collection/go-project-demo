package handleFile

import (
	"fmt"
	"go-project-demo/packages/pro3_proxypool/pkg/models"
	"go-project-demo/packages/pro3_proxypool/pkg/storage"
	"sync"
)

func getUsedIP(ip *models.IP, ipList chan<- *models.IP, wg *sync.WaitGroup) {
	used := storage.CheckIP(ip)
	if used {
		ipList <- ip
	}

	fmt.Println("处理 ip 中: ", ip.Data)

	wg.Done()
}

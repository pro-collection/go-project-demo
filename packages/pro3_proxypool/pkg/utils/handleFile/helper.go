package handleFile

import (
	"fmt"
	"go-project-demo/packages/pro3_proxypool/pkg/models"
	"go-project-demo/packages/pro3_proxypool/pkg/storage"
	"sync"
	"time"
)

func getUsedIP(ip *models.IP, ipList chan<- *models.IP, wg *sync.WaitGroup) {
	used := storage.CheckIP(ip)
	if used {
		ip.UpdateTime = time.Now()
		ipList <- ip
	}

	fmt.Println("处理 ip 是否可用: ", ip.Data, "  -- 结论： ", used)

	wg.Done()
}

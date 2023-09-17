package api

import (
	"encoding/json"
	"fmt"
	"go-project-demo/packages/pro3_proxypool/pkg/logger"
	"go-project-demo/packages/pro3_proxypool/pkg/models"
	"go-project-demo/packages/pro3_proxypool/pkg/storage"
	"go-project-demo/packages/pro3_proxypool/pkg/utils"
	"go-project-demo/packages/pro3_proxypool/pkg/utils/handleFile"
	"net/http"
	"os"
)

// VERSION for this program

// RunWithLocal 启动服务
func RunWithLocal() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ip", proxyHandler)
	//mux.HandleFunc("/https", findHandler)
	logger.Info(&logger.Params{
		Key:      logger.Key.BaseInfo,
		ModeName: "api",
		FuncName: "Run",
		Content:  "starting server: " + "127.0.0.1" + ":" + "3000",
	})
	http.ListenAndServe("127.0.0.1"+":"+"3000/ip", mux)
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("content-type", "application/json")
		file, err := os.OpenFile("ip.json", os.O_RDWR, 0644)
		fileContent, err := handleFile.ReadFile(file)
		if err != nil {
			fmt.Println("打开文件失败 - 确实文件 ip.json", err)
			panic(err)
		}

		var ipList []*models.IP
		err = json.Unmarshal(fileContent, &ipList)

		randomNum := utils.RandInt(0, len(ipList))

		b, err := json.Marshal(ipList[randomNum])

		if err != nil {
			return
		}

		_, err = w.Write(b)
		if err != nil {
			return
		}
	}
}

func findHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("content-type", "application/json")
		b, err := json.Marshal(storage.ProxyRandom())
		if err != nil {
			return
		}
		w.Write(b)
	}
}

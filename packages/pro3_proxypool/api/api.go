package api

import (
	"encoding/json"
	"go-project-demo/packages/pro3_proxypool/pkg/logger"
	"go-project-demo/packages/pro3_proxypool/pkg/setting"
	"go-project-demo/packages/pro3_proxypool/pkg/storage"
	"net/http"
)

// VERSION for this program
const VERSION = "/v2"

// Run 启动服务
func Run() {
	mux := http.NewServeMux()

	mux.HandleFunc(VERSION+"/pi", ProxyHandler)
	mux.HandleFunc(VERSION+"/https", FindHandler)
	logger.Info(&logger.Params{
		Key:      logger.Key.BaseInfo,
		ModeName: "api",
		FuncName: "Run",
		Content:  "starting server: " + setting.AppAddr + ":" + setting.AppPort,
	})
	http.ListenAndServe(setting.AppAddr+":"+setting.AppPort, mux)
}

func ProxyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("content-type", "application/json")
		b, err := json.Marshal(storage.ProxyRandom())
		if err != nil {
			return
		}

		_, err = w.Write(b)
		if err != nil {
			return
		}
	}
}

func FindHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("content-type", "application/json")

	}
}

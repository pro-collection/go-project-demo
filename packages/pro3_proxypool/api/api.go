package api

import (
	"encoding/json"
	"go-project-demo/packages/pro3_proxypool/pkg/storage"
	"net/http"
)

// VERSION for this program
const VERSION = "/v2"

func Run() {
	//mux := http.NewServeMux()

	//mux.Handler(VERSION + "/pi", )
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

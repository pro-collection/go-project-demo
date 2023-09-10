package api

import (
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
		//b, err := json.Marshal()
	}
}

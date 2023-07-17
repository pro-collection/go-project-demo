package main

import (
	"go-project-demo/packages/pro2/internal/routers"
	"net/http"
	"time"
)

func main() {
	// r := gin.Default()
	router := routers.NewRouter()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	_ = s.ListenAndServe()

	//r.GET("ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{"message": "hello world - yanle"})
	//})
	//
	//_ = r.Run()
}

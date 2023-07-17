package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-project-demo/packages/pro2/global"
	"go-project-demo/packages/pro2/internal/routers"
	setting2 "go-project-demo/packages/pro2/pkg/setting"
	"log"
	"net/http"
	"time"
)

// 初始化配置信息
func setupSetting() error {
	setting, err := setting2.NewSetting()
	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

func init() {
	err := setupSetting()
	if err != nil {
		log.Printf("init.setupSetting err %v", err)
	}
}

func main() {
	// 设置启动模式
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()

	log.Println("yanle - port: ", fmt.Sprintf(":%s", global.ServerSetting.HttpPort))

	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", global.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	_ = s.ListenAndServe()

	//r.GET("ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{"message": "hello world - yanle"})
	//})
	//
	//_ = r.Run()
}

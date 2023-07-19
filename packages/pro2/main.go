package main

import (
	"fmt"
	"go-project-demo/packages/pro2/global"
	"go-project-demo/packages/pro2/internal/model"
	"go-project-demo/packages/pro2/pkg/logger"
	setting2 "go-project-demo/packages/pro2/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
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

	err = setting.ReadSectionWithSecret("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

// 初始化数据库连接
func setupDBEngin() error {
	var err error
	global.DBEngin, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupLogger() error {
	appSetting := global.AppSetting
	fileName := fmt.Sprintf("%s/%s%s", appSetting.UploadSavePath, appSetting.LogFileName, appSetting.LogFileExt)

	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

func init() {
	err := setupSetting()
	if err != nil {
		log.Printf("init.setupSetting err %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Printf("init.setupLogger err %v", err)
	}
}

func main() {
	// 设置启动模式
	//gin.SetMode(global.ServerSetting.RunMode)
	//router := routers.NewRouter()
	//
	//log.Println("yanle - info: ", global.DatabaseSetting)
	//
	//s := &http.Server{
	//	Addr:           fmt.Sprintf(":%s", global.ServerSetting.HttpPort),
	//	Handler:        router,
	//	ReadTimeout:    global.ServerSetting.ReadTimeout,
	//	WriteTimeout:   global.ServerSetting.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//
	//_ = s.ListenAndServe()

	//global.Logger.Infof("%s: go/%s", "yanle", "base")

	//r.GET("ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{"message": "hello world - yanle"})
	//})
	//
	//_ = r.Run()
}

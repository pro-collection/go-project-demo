package initial

import (
	"go-project-demo/packages/pro3_proxypool/pkg/setting"
)

func GlobalInit() {
	//setting.NewContext()
	setting.NewLogService()

	//models.LoadDataBaseInfo()
	//
	//if setting.InstallLock {
	//	if err := models.NewEngine(); err != nil {
	//		// 日志记录
	//		logger.Fatal(&logger.Params{
	//			Key:      logger.Key.InitORMEnginError,
	//			ModeName: "initial",
	//			FuncName: "GlobalInit",
	//			Content:  fmt.Sprintf("fail to initialize ORM engin: %v", err),
	//		})
	//	}
	//
	//	models.HasEngin = true
	//}
	//
	//if models.EnableSQLite3 {
	//	logger.Info(&logger.Params{
	//		Key:      logger.Key.BaseInfo,
	//		ModeName: "initial",
	//		FuncName: "GlobalInit",
	//		Content:  "SQLite3 Supported",
	//	})
	//}
	//
	//if !setting.InstallLock {
	//	models.SetDataBaseInfo()
	//}
}

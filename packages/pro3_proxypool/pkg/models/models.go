package models

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/go-xorm/xorm"
	"go-project-demo/packages/pro3_proxypool/pkg/logger"
	"go-project-demo/packages/pro3_proxypool/pkg/setting"
	"go-project-demo/packages/pro3_proxypool/pkg/utils/fileTool"
	"os"
	"path/filepath"
	"xorm.io/core"
)

var (
	x             *xorm.Engine
	tables        []interface{}
	HasEngin      bool
	DBConfig      DBConfigStruct
	EnableSQLite3 bool
)

func init() {
	tables = append(tables, new(IP))
	gonicNames := []string{"SSL"}

	for _, name := range gonicNames {
		core.LintGonicMapper[name] = true
	}
}

// SetDataBaseInfo 从 ini 中获取数据
func SetDataBaseInfo() {
	var loggerParams = &logger.Params{
		ModeName: "setting",
		FuncName: "SetDataBaseInfo",
	}

	var x *xorm.Engine
	if err := NewTextEngin(x); err != nil {
		loggerParams.Key = logger.Key.FatalInfo
		loggerParams.Content = fmt.Sprintf("fail to set test ORM engin: %v", err)
		logger.Fatal(loggerParams)
	}

	config := ini.Empty()

	if fileTool.IsFile(setting.ConfFile) {
		if err := config.Append(setting.ConfFile); err != nil {
			loggerParams.Key = logger.Key.ErrorInfo
			loggerParams.Content = fmt.Sprintf("Fail to load conf '%s': %v", setting.ConfFile, err)
			logger.Error(loggerParams)
		}
	}

	config.Section("").Key("APP_NAME").SetValue(setting.AppName)

	// Save server config
	config.Section("server").Key("HTTP_ADDR").SetValue(setting.AppAddr)
	config.Section("server").Key("HTTP_PORT").SetValue(setting.AppPort)
	config.Section("server").Key("SESSION_EXPIRES").SetValue(setting.SessionExpires.String())

	// Save database config
	config.Section("database").Key("DB_TYPE").SetValue(DBConfig.Type)
	config.Section("database").Key("HOST").SetValue(DBConfig.Host)
	config.Section("database").Key("NAME").SetValue(DBConfig.Name)
	config.Section("database").Key("USER").SetValue(DBConfig.User)
	config.Section("database").Key("PASSWD").SetValue(DBConfig.Password)
	config.Section("database").Key("SSL_MODE").SetValue(DBConfig.SSLMode)
	config.Section("database").Key("PATH").SetValue(DBConfig.Path)

	// Change Installock value to true
	config.Section("security").Key("INSTALL_LOCK").SetValue("true")

	// Save log config
	config.Section("log").Key("MODE").SetValue("file")
	config.Section("log").Key("LEVEL").SetValue("Info")
	config.Section("log").Key("BUFFER_LEN").SetValue("100")
	config.Section("log").Key("ROOT_PATH").SetValue(setting.LogRootPath)

	os.MkdirAll(filepath.Dir(setting.ConfFile), os.ModePerm)

	if err := config.SaveTo(setting.ConfFile); err != nil {
		loggerParams.Key = logger.Key.FatalInfo
		loggerParams.Content = fmt.Sprintf("[Initial]Save config failed: %v", err)
		logger.Fatal(loggerParams)
	}

	loggerParams.Key = logger.Key.BaseInfo
	loggerParams.Content = "[Initial]Initialize database completed."
	logger.Info(loggerParams)
}

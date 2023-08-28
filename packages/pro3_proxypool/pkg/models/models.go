package models

import (
	"github.com/go-xorm/xorm"
	"go-project-demo/packages/pro3_proxypool/pkg/consts"
	"go-project-demo/packages/pro3_proxypool/pkg/setting"
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

//func SetEngin() (err error) {
//	//x, err = geten
//}

// LoadDataBaseInfo .
func LoadDataBaseInfo() {
	sec := setting.Config.Section("database")
	DBConfig.Type = sec.Key("DB_TYPE").String()

	// 设置使用 db 的类型
	switch DBConfig.Type {
	case consts.DBType.Sqlite3:
		setting.UseSQLite3 = true
		EnableSQLite3 = true
	case consts.DBType.Mysql:
		setting.UseMySQL = true
	case consts.DBType.Postgres:
		setting.UsePostgreSQL = true
	case consts.DBType.Mssql:
		setting.UseMSSQL = true
	}

	DBConfig.Host = sec.Key("HOST").String()
	DBConfig.Name = sec.Key("NAME").String()
	DBConfig.User = sec.Key("USER").String()
	if len(DBConfig.Password) == 0 {
		DBConfig.Password = sec.Key("PASSWD").String()
	}
	DBConfig.SSLMode = sec.Key("SSL_MODE").String()
	DBConfig.Path = sec.Key("PATH").MustString("data/ProxyPool.db")
}

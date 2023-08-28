package consts

import "unknwon.dev/clog/v2"

// 存放敞亮

var ValidLevels = []string{"trace", "info", "warn", "error", "fatal"}

var LevelNames = map[string]clog.Level{
	"trace": clog.LevelTrace,
	"info":  clog.LevelInfo,
	"warn":  clog.LevelWarn,
	"error": clog.LevelError,
	"fatal": clog.LevelFatal,
}

var EnvMode = envModeStruct{
	Dev:  "dev",
	Prod: "prod",
}

var DBType = dbTypeStruct{
	Sqlite3:  "sqlite3",
	Mysql:    "mysql",
	Postgres: "postgres",
	Mssql:    "mssql",
}

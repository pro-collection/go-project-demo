package models

import (
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"go-project-demo/packages/pro3_proxypool/pkg/consts"
	"go-project-demo/packages/pro3_proxypool/pkg/setting"
	"net/url"
	"os"
	"path"
	"strings"
	"unknwon.dev/clog/v2"
	"xorm.io/core"
)

func parseHostPort(host string) (ip string, port string) {
	ip = ""
	port = ""

	if strings.Contains(host, ":") {
		ip = strings.Split(host, ":")[0]
		port = strings.Split(host, ":")[1]
	}
	return ip, port
}

// todo yanlele getEngine 初始化 xorm 引擎
func getEngine() (*xorm.Engine, error) {
	connStr := ""
	param := "?"

	if strings.Contains(DBConfig.Name, param) {
		param = "&"
	}

	switch DBConfig.Type {
	case consts.DBType.Mysql:
		if DBConfig.Host[0] == '/' { // looks like a unix socket
			connStr = fmt.Sprintf(
				"%s:%s@unix(%s)/%s%scharset=utf8&parseTime=true",
				DBConfig.User,
				DBConfig.Password,
				DBConfig.Host,
				DBConfig.Name,
				param,
			)
		} else {
			connStr = fmt.Sprintf(
				"%s:%s@tcp(%s)/%s%scharset=utf8&parseTime=true",
				DBConfig.User,
				DBConfig.Password,
				DBConfig.Host,
				DBConfig.Name,
				param,
			)
		}
	case "postgres":
		host, port := parseHostPort(DBConfig.Host)
		if host[0] == '/' { // looks like a unix socket
			connStr = fmt.Sprintf("postgres://%s:%s@:%s/%s%ssslmode=%s&host=%s",
				url.QueryEscape(DBConfig.User), url.QueryEscape(DBConfig.Password), port, DBConfig.Name, param, DBConfig.SSLMode, host)
		} else {
			connStr = fmt.Sprintf("postgres://%s:%s@%s:%s/%s%ssslmode=%s",
				url.QueryEscape(DBConfig.User), url.QueryEscape(DBConfig.Password), host, port, DBConfig.Name, param, DBConfig.SSLMode)
		}
	case "mssql":
		host, port := parseHostPort(DBConfig.Host)
		connStr = fmt.Sprintf("server=%s; port=%s; database=%s; user id=%s; password=%s;", host, port, DBConfig.Name, DBConfig.User, DBConfig.Password)
	case "sqlite3":
		if !EnableSQLite3 {
			return nil, errors.New("this binary version does not build support for SQLite3")
		}
		if err := os.MkdirAll(path.Dir(DBConfig.Path), os.ModePerm); err != nil {
			return nil, fmt.Errorf("fail to create directories: %v", err)
		}
		connStr = "file:" + DBConfig.Path + "?cache=shared&mode=rwc"
	default:
		return nil, fmt.Errorf("unknown database type: %s", DBConfig.Type)
	}

	return xorm.NewEngine(DBConfig.Type, connStr)
}

func NewTextEngin(x *xorm.Engine) (err error) {
	x, err = getEngine()

	if err != nil {
		return fmt.Errorf("connect to database : %v", err)
	}

	x.SetMapper(core.GonicMapper{})

	return x.StoreEngine("InnoDB").Sync2(tables...)
}

func SetEngine() (err error) {
	x, err = getEngine()
	if err != nil {
		return fmt.Errorf("fail to connect to database: %v", err)
	}

	x.SetMapper(core.GonicMapper{})

	// todo yanlele 这个可能什么都获取不到
	sec := setting.Config.Section("log.xorm")
	logger, err := clog.NewFileWriter(
		path.Join(setting.LogRootPath, "xorm.log"),
		clog.FileRotationConfig{
			Rotate:  sec.Key("ROTATE").MustBool(true),
			Daily:   sec.Key("ROTATE_DAILY").MustBool(true),
			MaxSize: sec.Key("MAX_SIZE").MustInt64(100) * 1024 * 1024,
			MaxDays: sec.Key("MAX_DAYS").MustInt64(3),
		},
	)

	if err != nil {
		return fmt.Errorf("fail to create 'xorm.log': %v", err)
	}

	if !setting.DebugMode {
		x.SetLogger(xorm.NewSimpleLogger3(logger, xorm.DEFAULT_LOG_PREFIX, xorm.DEFAULT_LOG_FLAG, core.LOG_WARNING))
	} else {
		x.SetLogger(xorm.NewSimpleLogger(logger))
	}

	x.ShowSQL(true)

	return nil
}

func NewEngine() (err error) {
	if err = SetEngine(); err != nil {
		return err
	}

	if err = x.StoreEngine("InnoDB").Sync2(tables...); err != nil {
		return fmt.Errorf("sysnc databse struct error: %v", err)
	}

	return nil
}

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

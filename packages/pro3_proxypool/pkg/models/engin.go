package models

import (
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"go-project-demo/packages/pro3_proxypool/pkg/consts"
	"net/url"
	"os"
	"path"
	"strings"
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

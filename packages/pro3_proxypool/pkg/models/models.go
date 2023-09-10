package models

import (
	"github.com/go-xorm/xorm"
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

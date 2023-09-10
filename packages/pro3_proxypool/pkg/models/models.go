package models

import (
	"github.com/go-xorm/xorm"
	"time"
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

func getAll() ([]*IP, error) {
	tmpIp := make([]*IP, 0)
	err := x.Where("speed <= 1000").Find(&tmpIp)

	if err != nil {
		return nil, err
	}

	return tmpIp, nil
}

func GetAll() ([]*IP, error) {
	return getAll()
}

// NewIp 默认 ip
func NewIp() *IP {
	return &IP{
		Speed:      100,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
}

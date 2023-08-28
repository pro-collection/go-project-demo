package models

import (
	"time"
)

type IP struct {
	ID         int64     `xorm:"pk autoincr" json:"-"`
	Data       string    `xorm:"NOT NULL unique" json:"ip"`
	Type1      string    `xorm:"NOT NULL" json:"type1"`
	Type2      string    `xorm:"NULL" json:"type2,omitempty"`
	Speed      int64     `xorm:"NOT NULL" json:"speed,omitempty"`  //连接速度
	Source     string    `xorm:"NOT NULL" json:"source,omitempty"` //代理来源
	CreateTime time.Time `xorm:"NOT NULL" json:"-"`
	UpdateTime time.Time `xorm:"NOT NULL" json:"-"`
}

type DBConfigStruct struct {
	Type, Host, Name, User, Password, Path, SSLMode string
}

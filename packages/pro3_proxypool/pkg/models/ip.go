package models

import (
	"go-project-demo/packages/pro3_proxypool/pkg/logger"
	"time"
)

// TestHttps xorm 语法要研究一下
func TestHttps() bool {
	has, err := x.Exist(&IP{Type1: "https"})
	if err != nil {
		return false
	}

	return has
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

func FindAll(value string) ([]*IP, error) {
	tempIp := make([]*IP, 0)
	switch value {
	case "http":
		err := x.Where("speed <= 1000 and type1=?", value).Find(&tempIp)
		if err != nil {
			return tempIp, err
		}
	case "https":
		hasHttps := TestHttps()
		if hasHttps == false {
			return tempIp, nil
		}
		err := x.Where("speed <= 1000 and type1=?", value).Find(&tempIp)
		if err != nil {
			logger.Error(&logger.Params{
				Key:      logger.Key.ErrorInfo,
				ModeName: "models",
				FuncName: "FindAll",
				Content:  err.Error(),
			})

			return tempIp, err
		}
	default:
		return tempIp, nil
	}

	return tempIp, nil
}

func CountIps() int64 {
	count, _ := x.Where("id>=?", 0).Count(new(IP))
	return count
}

func InsertIps(ip *IP) (err error) {
	ses := x.NewSession()

	defer ses.Close()

	if err := ses.Begin(); err != nil {
		return err
	}

	if _, err = ses.Insert(ip); err != nil {
		return err
	}

	return ses.Commit()
}

func Update(ip *IP) error {
	temp := ip
	temp.UpdateTime = time.Now()
	_, err := x.Id(1).Update(temp)

	if err != nil {
		return err
	}

	return nil
}

func DeleteIP(ip *IP) error {
	_, err := x.Delete(ip)
	if err != nil {
		return err
	}

	return nil
}

func NewIP() *IP {
	return &IP{
		Speed:      -1,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
}

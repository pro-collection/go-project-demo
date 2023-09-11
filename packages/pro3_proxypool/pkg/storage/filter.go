package storage

import (
	"fmt"
	"go-project-demo/packages/pro3_proxypool/pkg/logger"
	"go-project-demo/packages/pro3_proxypool/pkg/models"
	"go-project-demo/packages/pro3_proxypool/pkg/utils"
	"sync"
)

func ProxyRandom() (ip *models.IP) {
	ips, err := models.GetAll()

	ip = &models.IP{}

	x := len(ips)

	var loggerParams = &logger.Params{
		Key:      logger.Key.BaseInfo,
		ModeName: "storage",
		FuncName: "ProxyRandom",
	}

	loggerParams.Content = fmt.Sprintf("len(ips) = %d", x)
	logger.Info(loggerParams)

	if x == 0 {
		loggerParams.Key = logger.Key.WarnInfo
		loggerParams.Content = "no ips"
		logger.Warn(loggerParams)
		return ip
	}

	if err != nil {
		loggerParams.Key = logger.Key.WarnInfo
		loggerParams.Content = err.Error()
		logger.Warn(loggerParams)
		return ip
	}

	randomNum := utils.RandInt(0, x)

	return ips[randomNum]
}

func ProxyFind(value string) (ip *models.IP) {
	ips, err := models.FindAll(value)
	if err != nil {
		logger.Warn(&logger.Params{
			Key:      logger.Key.WarnInfo,
			ModeName: "storage",
			FuncName: "ProxyFind",
			Content:  err.Error(),
		})
		return models.NewIp()
	}

	x := len(ips)
	if x == 0 {
		logger.Warn(&logger.Params{
			Key:      logger.Key.WarnInfo,
			ModeName: "storage",
			FuncName: "ProxyFind",
			Content:  err.Error(),
		})
		return models.NewIp()
	}

	randomNum := utils.RandInt(0, x)
	return ips[randomNum]
}

func CheckProxyDB() {
	loggerParams := &logger.Params{
		Key:      logger.Key.BaseInfo,
		ModeName: "storage",
		FuncName: "CheckProxyDB",
	}

	x := models.CountIps()
	loggerParams.Content = fmt.Sprintf("Before check, DB has: %d records.", x)
	logger.Info(loggerParams)

	ips, err := models.GetAll()

	if err != nil {
		loggerParams.Key = logger.Key.WarnInfo
		loggerParams.Content = err.Error()
		logger.Warn(loggerParams)
	}

	var wg sync.WaitGroup

	for _, value := range ips {
		wg.Add(1)
		go func(v *models.IP) {
			// todo yanlele CheckIP
			// todo yanlele ProxyDel
		}(value)
	}

	wg.Wait()

	x = models.CountIps()

	loggerParams.Key = logger.Key.BaseInfo
	loggerParams.Content = fmt.Sprintf("After check, DB has: %d records.", x)
	logger.Info(loggerParams)
}

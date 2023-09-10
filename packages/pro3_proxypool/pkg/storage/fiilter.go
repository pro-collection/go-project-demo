package storage

import (
	"fmt"
	"go-project-demo/packages/pro3_proxypool/pkg/logger"
	"go-project-demo/packages/pro3_proxypool/pkg/models"
	"go-project-demo/packages/pro3_proxypool/pkg/utils"
)

func ProxyRandom() (ip *models.IP) {
	ips, err := models.GetAll()

	x := len(ips)

	var loggerParams = logger.Params{
		Key:      logger.Key.BaseInfo,
		ModeName: "storage",
		FuncName: "ProxyRandom",
	}

	loggerParams.Content = fmt.Sprintf("len(ips) = %d", x)
	logger.Info(loggerParams)

	if err != nil || x == 0 {
		loggerParams.Key = logger.Key.WarnInfo
		loggerParams.Content = err.Error()
		logger.Warn(loggerParams)
	}

	randomNum := utils.RandInt(0, x)

	return ips[randomNum]
}

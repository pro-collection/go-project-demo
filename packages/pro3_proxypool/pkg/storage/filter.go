package storage

import (
	"crypto/tls"
	"fmt"
	"go-project-demo/packages/pro3_proxypool/pkg/logger"
	"go-project-demo/packages/pro3_proxypool/pkg/models"
	"go-project-demo/packages/pro3_proxypool/pkg/utils"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
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
			if !CheckIP(v) {
				ProxyDelete(v)
			}
			wg.Done()
		}(value)
	}

	wg.Wait()

	x = models.CountIps()

	loggerParams.Key = logger.Key.BaseInfo
	loggerParams.Content = fmt.Sprintf("After check, DB has: %d records.", x)
	logger.Info(loggerParams)
}

func ProxyAdd(ip *models.IP) {
	_ = models.InsertIps(ip)
}

func ProxyDelete(ip *models.IP) {
	_ = models.DeleteIP(ip)
}

func CheckIP(ip *models.IP) bool {
	var pollURL string
	var testIP string

	if ip.Type2 == "https" {
		testIP = "https://" + ip.Data
		pollURL = "https://httpbin.org/get?show_env=1"
	} else {
		testIP = "http://" + ip.Data
		pollURL = "http://httpbin.org/get?show_env=1"
	}

	proxy, _ := url.Parse(testIP)

	logger.Info(&logger.Params{
		Key:      logger.Key.BaseInfo,
		ModeName: "storage",
		FuncName: "CheckIP",
		Content:  testIP,
	})

	begin := time.Now()

	tlsConfig := &tls.Config{InsecureSkipVerify: true}

	netTransport := &http.Transport{
		Proxy:               http.ProxyURL(proxy),
		TLSClientConfig:     tlsConfig,
		MaxIdleConnsPerHost: 50,
	}

	httpClient := &http.Client{
		Timeout:   time.Second * 20,
		Transport: netTransport,
	}

	request, _ := http.NewRequest("GET", pollURL, nil)

	request.Header.Add("accept", "text/plain")

	resp, err := httpClient.Do(request)

	if err != nil {
		logger.Warn(&logger.Params{
			Key:      logger.Key.WarnInfo,
			ModeName: "storage",
			FuncName: "CheckIP",
			Content:  fmt.Sprintf("[CheckIP] testIP = %s, pollURL = %s: Error = %v", testIP, pollURL, err),
		})

		return false
	}

	// 终止函数执行的时候， 退出请求
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode == 200 {
		// todo yanlele 判断返回数据的合法性

		ip.Speed = time.Now().Sub(begin).Milliseconds()

		// 保存更新
		// todo yanlele 保存更新
		if err = models.Update(ip); err != nil {
			logger.Warn(&logger.Params{
				Key:      logger.Key.WarnInfo,
				ModeName: "storage",
				FuncName: "CheckIP",
				Content:  fmt.Sprintf("[CheckIP] Update IP = %v Error = %v", *ip, err),
			})
		}
		return true
	}

	return false
}

func CheckProxy(ip *models.IP) {
	if CheckIP(ip) {
		ProxyAdd(ip)
	}
}

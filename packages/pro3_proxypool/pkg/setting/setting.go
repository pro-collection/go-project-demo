package setting

import (
	"fmt"
	"github.com/go-ini/ini"
	"go-project-demo/packages/pro3_proxypool/pkg/consts"
	"go-project-demo/packages/pro3_proxypool/pkg/utils"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
	"unknwon.dev/clog/v2"
)

var (
	/*
		App settings
	*/
	AppVer  string
	AppName string
	AppURL  string
	AppPath string
	AppAddr string
	AppPort string
	DevMode string

	/*
		Global setting objects
	*/
	Config    *ini.File
	DebugMode bool
	IsWindows bool
	ConfFile  string

	/*
		Database settings
	*/
	UseSQLite3    bool
	UseMySQL      bool
	UsePostgreSQL bool
	UseMSSQL      bool

	/*
		Log settings
	*/
	LogRootPath string
	LogModes    []string
	LogConfigs  []interface{}

	/*
		Security settings
	*/
	InstallLock bool // true mean installed

	/*
		OAuth2 settings
	*/
	SessionExpires time.Duration
)

// execPath
// 获取可执行的 path
// 如果出现了位置信息读取失败的情况， 直接终止程序即可
func execPath() string {
	exePath, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get executable path:", err)
		panic(err)
	}

	exeDir := filepath.Dir(exePath)
	absExeDir, err := filepath.Abs(exeDir)
	if err != nil {
		fmt.Println("Failed to get absolute executable directory:", err)
		panic(err)
	}

	return absExeDir
}

func init() {
	IsWindows = runtime.GOOS == "windows"

	env, _ := utils.GetEnv("env")

	if env == consts.EnvMode.Dev {
		absExeDir := execPath()
		AppPath = absExeDir + "pro3_proxypool"
	} else {
		//这个场景是提供给外部使用场景
		AppPath = execPath()
	}
}

// WorkDir 获取工作区地址
func WorkDir() (string, error) {
	wd := os.Getenv("ALIGN_WORK_DIR")
	if len(wd) > 0 {
		return wd, nil
	}

	i := strings.LastIndex(AppPath, "/")
	if i == -1 {
		return AppPath, nil
	}

	return AppPath[:i], nil
}

func NewContext() {
	// todo yanlele 暂时不知道这个是干啥的
	//workdir, err := WorkDir()
	//if err != nil {
	//	clog.Fatal("Fail to get work directory: %v", err)
	//}

	var err error

	// 获取配置文件地址
	ConfFile = path.Join(AppPath, "conf/app.ini")

	// 读取配置
	Config, err = ini.Load(ConfFile)

	if err != nil {
		fmt.Println(fmt.Sprintf("Fail to parse %s: %v", ConfFile, err))
		panic(err)
	}

	// 配置名称转换， 全部都转换诶大写
	// 具体可以参考这篇文档： https://github.com/yanlele/golang-index/issues/4
	Config.NameMapper = ini.SnackCase

	// Load security config
	InstallLock = Config.Section("security").Key("INSTALL_LOCK").MustBool(false)

	// Load server config
	sec := Config.Section("server")
	AppName = Config.Section("").Key("APP_NAME").MustString("ProxyPool")
	AppURL = sec.Key("ROOT_URL").MustString("http://localhost:3000/")
	if AppURL[len(AppURL)-1] != '/' {
		AppURL += "/"
	}
	AppAddr = sec.Key("HTTP_ADDR").MustString("0.0.0.0")
	AppPort = sec.Key("HTTP_PORT").MustString("3001")
	SessionExpires = sec.Key("SESSION_EXPIRES").MustDuration(time.Hour * 24 * 7)

	// 获取开发环境
	DevMode = Config.Section("").Key("MODE").MustString(consts.EnvMode.Dev)
}

func validateFunc(v string) string {
	if utils.IncludesWithString(consts.ValidLevels, v) {
		return v
	}
	return "trice"
}

// NewLogService 初始化日志
func NewLogService() {
	err := clog.NewConsole(
		//100,
		clog.ConsoleConfig{
			Level: clog.LevelInfo,
		},
	)
	if err != nil {
		clog.Warn("unable to create new logger: " + err.Error())
	}
}

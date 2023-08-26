package setting

import (
	"github.com/go-ini/ini"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	clog "unknwon.dev/clog/v2"
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

	/*
		Global setting objects
	*/
	Cfg       *ini.File
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
func execPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}

	return filepath.Abs(file)
}

func init() {
	IsWindows = runtime.GOOS == "windows"
	var err error
	AppPath, err = execPath()
	if err != nil {
		clog.Fatal("Fail to get app path: %v\n", err)
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
	workdir, err := WorkDir()
	if err != nil {
		clog.Fatal("Fail to get work directory: %v", err)
	}

	// 获取配置文件地址
	ConfFile = path.Join(workdir, "conf/app.ini")

	// 读取配置
	Cfg, err = ini.Load(ConfFile)

	if err != nil {
		clog.Fatal("Fail to parse %s: %v", ConfFile, err)
	}

	// 配置名称转换， 全部都转换诶大写
	// 具体可以参考这篇文档： https://github.com/yanlele/golang-index/issues/4
	Cfg.NameMapper = ini.SnackCase

	// Load security config
	InstallLock = Cfg.Section("security").Key("INSTALL_LOCK").MustBool(false)

	// Load server config
	sec := Cfg.Section("server")
	AppName = Cfg.Section("").Key("APP_NAME").MustString("ProxyPool")
	AppURL = sec.Key("ROOT_URL").MustString("http://localhost:3000/")
	if AppURL[len(AppURL)-1] != '/' {
		AppURL += "/"
	}
	AppAddr = sec.Key("HTTP_ADDR").MustString("0.0.0.0")
	AppPort = sec.Key("HTTP_PORT").MustString("3001")
	SessionExpires = sec.Key("SESSION_EXPIRES").MustDuration(time.Hour * 24 * 7)
}

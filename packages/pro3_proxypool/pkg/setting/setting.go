package setting

import (
	"fmt"
	"github.com/go-ini/ini"
	"go-project-demo/packages/pro3_proxypool/pkg/consts"
	"go-project-demo/packages/pro3_proxypool/pkg/logger"
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
func execPath() (string, error) {
	file, err := os.Getwd()
	if err != nil {
		return "", err
	}

	path, err := filepath.Abs(file)
	return path, err
}

func init() {
	IsWindows = runtime.GOOS == "windows"

	args := os.Args
	var env string

	for index, arg := range args {
		if index == 1 {
			env = arg
		}
	}

	var err error

	if env == "dev" {
		//var err error
		exePath, err := os.Getwd()
		if err != nil {
			fmt.Println("Failed to get executable path:", err)
			return
		}

		exeDir := filepath.Dir(exePath)
		absExeDir, err := filepath.Abs(exeDir)
		if err != nil {
			fmt.Println("Failed to get absolute executable directory:", err)
			return
		}

		AppPath = absExeDir + "/go-project-demo/packages/pro3_proxypool"
	} else {
		//这个场景是提供给外部使用场景
		AppPath, err = execPath()
		if err != nil {
			fmt.Printf("Fail to get app path: %v\n", err)
		}
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
		//logger.Fatal(logger.Params{
		//	Key:      logger.Key.GetConfigFail,
		//	ModeName: "setting",
		//	FuncName: "NewContext",
		//	Content:  "初始化配置失败, 请检测文件路径是否存在, 请检测路径下配置是否合法",
		//	Extend:   fmt.Sprintf("Fail to parse %s: %v", ConfFile, err),
		//})
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
	// 通过开发环境，获取日志的级别
	if DebugMode {
		LogModes = strings.Split("console", ",")
	} else {
		LogModes = strings.Split(Config.Section("log").Key("mode").MustString("console"), ",")
	}

	for _, mode := range LogModes {
		mode = strings.ToLower(strings.TrimSpace(mode))
		currentMode := "log." + mode
		sec, err := Config.GetSection(currentMode)
		if err != nil {
			logger.Fatal(logger.Params{
				Key:      logger.Key.UnknownLoggerMode,
				ModeName: "setting",
				FuncName: "NewLogService",
				Content:  fmt.Sprintf("Unknown logger mode: %s", mode),
			})
		}

		name := Config.Section(currentMode).Key("LEVEL").Validate(validateFunc)

		// 日志级别
		level := consts.LevelNames[name]

		// 只支持一下两种模式， 可以自行扩展
		switch mode {
		case "console":
			//bufferSize := Config.Section("log").Key("BUFFER_LEN").MustInt64(10000)
			err = clog.NewConsole(
				//100,
				clog.ConsoleConfig{
					Level: level,
				},
			)
			if err != nil {
				clog.Warn("unable to create new logger: " + err.Error())
			}
			break
		case "file":
			// 日志写入到文件
			logPath := path.Join(LogRootPath, "ProxyPool.log")
			err = os.MkdirAll(path.Dir(logPath), os.ModePerm)
			if err != nil {
				clog.Warn("Fail to create log directory '%s': %v", path.Dir(logPath), err)
			}

			err = clog.NewFile(clog.FileConfig{
				Level:    level,
				Filename: logPath,
				FileRotationConfig: clog.FileRotationConfig{
					Rotate:   sec.Key("LOG_ROTATE").MustBool(true),
					Daily:    sec.Key("DAILY_ROTATE").MustBool(true),
					MaxSize:  1 << uint(sec.Key("MAX_SIZE_SHIFT").MustInt(28)),
					MaxLines: sec.Key("MAX_LINES").MustInt64(1000000),
					MaxDays:  sec.Key("MAX_DAYS").MustInt64(7),
				},
			})
		}

		logger.Info(logger.Params{
			Key:      logger.Key.LoggerMode,
			ModeName: "setting",
			FuncName: "NewLogService",
			Content:  fmt.Sprintf("Log Mode: %s (%s)", utils.GetTitle(mode), utils.GetTitle(name)),
		})
	}

	// Make sure everyone gets version info printed.
	logger.Info(logger.Params{
		Key:      logger.Key.AppInfo,
		ModeName: "setting",
		FuncName: "NewLogService",
		Content:  fmt.Sprintf("app_name: %s, app_version: %s", AppName, AppVer),
	})
}

package logger

// LoggerKeyStruct 定义而已
type loggerKeyStruct struct {
	Path              string
	UnknownLoggerMode string
	LoggerMode        string
	AppInfo           string
	GetConfigFail     string
	InitORMEnginError string
	BaseInfo          string
	FatalInfo         string
	ErrorInfo         string
	WarnInfo          string
}

var Key = loggerKeyStruct{
	Path:              "yanle_logger_path",
	UnknownLoggerMode: "unknown logger mode",
	LoggerMode:        "logger mode",
	AppInfo:           "app info",
	GetConfigFail:     "GetConfigFail",
	InitORMEnginError: "InitORMEnginError",
	BaseInfo:          "BaseInfo",
	FatalInfo:         "FatalInfo",
	ErrorInfo:         "ErrorInfo",
	WarnInfo:          "WarnInfo",
}

type Params struct {
	Key      string
	ModeName string
	FuncName string
	Content  string
	Error    error
	Extend   interface{}
}

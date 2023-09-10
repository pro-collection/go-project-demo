package logger

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
}

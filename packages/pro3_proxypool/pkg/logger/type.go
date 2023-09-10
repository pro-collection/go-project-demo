package logger

// LoggerKeyStruct 定义而已
type loggerKeyStruct struct {
	Path              string
	UnknownLoggerMode string
	LoggerMode        string
	AppInfo           string
	GetConfigFail     string
}

type Params struct {
	Key      string
	ModeName string
	FuncName string
	Content  string
	Error    error
	Extend   interface{}
}
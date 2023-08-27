package consts

// 不对外暴露， 定义而已
type loggerKey struct {
	Path string
}

var LoggerKey = loggerKey{
	Path: "yanle_logger_path",
}

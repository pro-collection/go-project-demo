package consts

type envModeStruct struct {
	Dev  string
	Prod string
}

// 不对外暴露， 定义而已
type loggerKey struct {
	Path string
}

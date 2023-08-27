package logger

import (
	"fmt"
	"unknwon.dev/clog/v2"
)

type Params struct {
	Key      string
	ModeName string
	FuncName string
	Content  string
	Error    error
}

func getParams(params Params) string {
	str := fmt.Sprintf("[%s] %s.%s - %s", params.Key, params.ModeName, params.FuncName, params.Content)
	return str
}

func Trace(params Params) {
	clog.Trace(getParams(params))
}

func Info(params Params) {
	clog.Info(getParams(params))
}

func Warn(params Params) {
	clog.Warn(getParams(params))
}

func Error(params Params) {
	clog.Error(getParams(params))
}

func Fatal(params Params) {
	clog.Fatal(getParams(params))
}

package logger

import (
	"fmt"
	"unknwon.dev/clog/v2"
)

func getParams(params Params) string {
	var str string

	if params.Extend != nil {
		str = fmt.Sprintf(
			"[%s] %s.%s - %s | extned info: %s",
			params.Key,
			params.ModeName,
			params.FuncName,
			params.Content,
			params.Extend,
		)

		return str
	}

	str = fmt.Sprintf(
		"[%s] %s.%s - %s",
		params.Key,
		params.ModeName,
		params.FuncName,
		params.Content,
	)

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

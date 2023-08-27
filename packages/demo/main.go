package main

import (
	"fmt"
	"unknwon.dev/clog/v2"
)

/*
err = clog.NewConsole(
				//100,
				clog.ConsoleConfig{
					Level: level,
				},
			)
			clog.Warn("yanle  测试 日志")
			if err != nil {
				clog.Warn("unable to create new logger: " + err.Error())
			}
*/

func main() {
	var bufferSize = 100
	err := clog.NewConsole(
		bufferSize,
		clog.ConsoleConfig{
			Level: clog.LevelTrace,
		},
	)
	clog.Warn("yanle  测试 日志")
	if err != nil {
		clog.Warn("unable to create new logger: " + err.Error())
	}
	if err != nil {
		fmt.Println("Failed to initialize logger:", err)
		return
	}

	defer clog.Stop()

	clog.Trace("This is a trace log")
}

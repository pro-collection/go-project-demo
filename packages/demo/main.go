package main

import (
	"context"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

type Logger interface {
	Debug(msg string, field map[string]interface{})
	Info(msg string, field map[string]interface{})
	Warn(msg string, field map[string]interface{})
	Error(msg string, field map[string]interface{})
	Fatal(msg string, field map[string]interface{})
}

func main() {

	log.SetOutput(&lumberjack.Logger{
		Filename:   "./foo.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28,    //days
		Compress:   false, // disabled by default
	})

	parentContext := context.Background()

	ctx := context.WithValue(parentContext, "key", "value")

	log.Println("demo", ctx.Value("key"))
}

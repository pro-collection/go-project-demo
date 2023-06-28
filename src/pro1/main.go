package main

import (
	"go-project-demo/src/pro1/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}

	//fmt.Println(timer.GetNowTime())
}

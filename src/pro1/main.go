package main

import (
	"go-project-demo/src/pro1/cmd"
	"log"
)

var name string

func main() {
	err := cmd.Execute()

	if err != nil {
		log.Fatalf("cmd.Execute err : %v\n", err)
	}
}

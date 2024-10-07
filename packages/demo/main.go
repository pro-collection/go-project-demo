package main

import (
	"fmt"
	"go-project-demo/packages/demo/demos"
)

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Name     string
	Age      int
	JobTitle string
}

func main() {

	MConfig, _ := demos.LoadJsonData("packages/demo/demos/config.json", "")

	fmt.Println("MConfig: ", *MConfig)

	fmt.Print("MConfig.UDPTimeout： ", MConfig.UDPTimeout)

	for index, user := range MConfig.UserGroups {
		fmt.Println("user: ", user)
		fmt.Println("index: ", index)
	}

}

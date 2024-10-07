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
	fmt.Printf(demos.Demo1())
}

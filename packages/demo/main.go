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
	body, err := demos.Demo1("https://juejin.cn/post/7311603432929984552")

	// 存在错误
	if err != nil {
		return
	}

	fmt.Println(body)
}

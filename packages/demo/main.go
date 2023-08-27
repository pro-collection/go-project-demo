package main

import (
	"fmt"
	"os"
)

func main() {
	env := os.Args[1]

	fmt.Println("程序名称:", env)

	//if len(args) > 1 {
	//	fmt.Println("启动参数:")
	//	for i, arg := range args[1:] {
	//		fmt.Printf("%d: %s\n", i+1, arg)
	//	}
	//} else {
	//	fmt.Println("没有启动参数.")
	//}
}

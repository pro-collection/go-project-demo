package main

import "fmt"

type paramsType = map[string]string

func main() {
	params := make(paramsType)

	params["name"] = "yanle"

	go func() {
		res := params["age"]
		fmt.Println(res)
	}()

}

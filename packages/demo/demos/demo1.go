package demos

import (
	"fmt"
	"io"
	"net/http"
)

/*
获取 url 链接的 dome 节点

参数： url

返回值： body error
*/
func Demo1(url string) (string, error) {
	// 发送 GET 请求
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	defer response.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return "", err
	}

	return string(body), nil
}

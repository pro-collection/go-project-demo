package handleFile

import (
	"encoding/json"
	"fmt"
	"go-project-demo/packages/pro3_proxypool/pkg/getter"
	"os"
)

func FindFile(filename string) (file *os.File, err error) {
	if _, err = os.Stat(filename); os.IsNotExist(err) {
		// 文件不存在，创建文件
		file, err = os.Create(filename)
		if err != nil {
			fmt.Println("创建文件失败:", err)
			return nil, err
		}

		fmt.Println("文件创建成功")
	} else {
		// 文件存在，直接使用
		file, err = os.OpenFile(filename, os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("打开文件失败:", err)
			return nil, err
		}

		fmt.Println("文件已存在，可以直接使用")
	}

	return
}

// WriteFile
// 写入文件
func WriteFile(file *os.File) {
	ipList := getter.IP89()
	jsonData, _ := json.Marshal(ipList)

	_, err := file.WriteString(string(jsonData) + "\n")
	if err != nil {
		fmt.Println("写入文件失败:", err)
		return
	}
}

// ReadFile
// 读取文件， 返回 []byte
func ReadFile(file *os.File) ([]byte, error) {
	// 获取文件大小
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("获取文件信息失败:", err)
		return nil, err
	}
	fileSize := fileInfo.Size()

	buffer := make([]byte, fileSize)
	// 读取文件内容
	_, err = file.Read(buffer)

	if err != nil {
		fmt.Println("读取文件失败:", err)
		return nil, err
	}

	return buffer, nil
}
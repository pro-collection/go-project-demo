package utils

import (
	"fmt"
	"os"
)

// IsFile
// 判断是否是文件
func IsFile(filePath string) bool {
	f, e := os.Stat(filePath)
	if e != nil {
		return false
	}

	return !f.IsDir()
}

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
		file, err = os.Open(filename)
		if err != nil {
			fmt.Println("打开文件失败:", err)
			return nil, err
		}

		fmt.Println("文件已存在，可以直接使用")
	}

	return
}

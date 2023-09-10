package utils

import "os"

// IsFile
// 判断是否是文件
func IsFile(filePath string) bool {
	f, e := os.Stat(filePath)
	if e != nil {
		return false
	}

	return !f.IsDir()
}

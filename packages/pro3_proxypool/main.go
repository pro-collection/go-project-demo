package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// execPath returns the executable path.
func execPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	return filepath.Abs(file)
}

func main() {
	filePath, err := execPath()

	if err != nil {
		return
	}

	fmt.Println("path： ", filePath)
	//fmt.Println("console", strings.Split("console", ","))

	// todo yanlele 虽然配置完成， 但是读取配置好像有点儿问题
	//fmt.Println("LogRootPath: ", setting.Cfg.Section("log").Key("ROOT_PATH").MustString(""))
}

package main

import (
	"go-project-demo/packages/pro3_proxypool/pkg/consts"
	"go-project-demo/packages/pro3_proxypool/pkg/initial"
	"go-project-demo/packages/pro3_proxypool/pkg/setting"
	"os"
	"os/exec"
	"path/filepath"
	"unknwon.dev/clog/v2"
)

// execPath returns the executable path.
func execPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	return filepath.Abs(file)
}

func deferExec() {
	defer clog.Stop()
}

func main() {
	// 初始化
	initial.GlobalInit()

	filePath, err := execPath()

	if err != nil {
		return
	}

	clog.Trace("path： ", filePath)
	//fmt.Println("console", strings.Split("console", ","))

	// todo yanlele 虽然配置完成， 但是读取配置好像有点儿问题
	clog.Trace(consts.LoggerKey.Path, setting.Cfg.Section("log").Key("ROOT_PATH").MustString(""))

	deferExec()
}

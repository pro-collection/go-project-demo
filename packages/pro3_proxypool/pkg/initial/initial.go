package initial

import (
	"go-project-demo/packages/pro3_proxypool/pkg/setting"
	"unknwon.dev/clog/v2"
)

func GlobalInit() {
	setting.NewContext()
	setting.NewLogService()

	clog.Trace("log path: %s", setting.LogRootPath)
}

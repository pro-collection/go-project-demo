package initial

import (
	"go-project-demo/packages/pro3_proxypool/pkg/models"
	"go-project-demo/packages/pro3_proxypool/pkg/setting"
)

func GlobalInit() {
	setting.NewContext()
	setting.NewLogService()

	models.LoadDataBaseInfo()

}

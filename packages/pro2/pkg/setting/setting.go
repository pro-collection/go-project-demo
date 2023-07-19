package setting

import "github.com/spf13/viper"

type Setting struct {
	// 用于存放公共的配置
	vp *viper.Viper

	// 读取一些非公开的配置
	vps *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vps := viper.New()

	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	vp.AddConfigPath("packages/pro2/configs")

	vps.SetConfigName("config_secret")
	vps.SetConfigType("yaml")
	vps.AddConfigPath("packages/pro2/configs")

	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = vps.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp, vps}, nil
}

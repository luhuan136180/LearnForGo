package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigFile("configs/config.yaml")
	vp.SetConfigType("yaml")     // SetConfigType设置远程源返回的配置类型。“json”。
	vp.AddConfigPath("configs/") //设置其配置路径为相对路径

	err := vp.ReadInConfig() // 查找并读取配置文件
	if err != nil {          // 处理读取配置文件的错误
		return nil, err
	}
	return &Setting{vp}, nil
}

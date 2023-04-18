package global

import (
	"github.com/go-programming-tour-book/blog-service/pkg/logger"
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
)

var (
	//全局声明，将配置信息和应用程序关联
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	//定义一个Logger对象便于项目使用
	Logger *logger.Logger //用于日志组件的初始化。
	//增加JWT的全局对象
	JWTSetting *setting.JWTSettingS
	//email对应的配置全局对象
	EmailSetting *setting.EmailSettingS
)

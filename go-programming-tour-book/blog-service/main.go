package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/model"
	"github.com/go-programming-tour-book/blog-service/internal/routers"
	"github.com/go-programming-tour-book/blog-service/pkg/logger"
	setting "github.com/go-programming-tour-book/blog-service/pkg/setting"
	"github.com/go-programming-tour-book/blog-service/pkg/tracer"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init() {
	//初始化viper结构体，将配置文件信息反序列化到对应的结构体中
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	//初始化自定义日志
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

	//一下代码：测试是否成功启用自定义logger
	//global.Logger.Infof("%s: go-programming-tour-book/%s", "eddycjy", "blog-service")

	//初始化mysql的链接
	err = setipDBEngnie()
	if err != nil {
		//写入日志
		log.Fatalf("init.setupDBEngine err:%v", err)
	}

	//	初始化Tracer，用于链路追踪
	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err :%v", err)
	}
}

// @title 博客系统
// @version 1.0
// @description Go 语言编程之旅：一起用 Go 做项目
// @termsOfService https://github.com/go-programming-tour-book
func main() {
	gin.SetMode(global.ServerSetting.RunMode) //设置gin的运行模式
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

//初始化配置文件对应的信息结构体
func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	//log.Println(global.ServerSetting)
	//log.Println(global.AppSetting)
	//log.Println(global.DatabaseSetting)

	//设置了 JWT 令牌的 Secret（密钥）为 eddycjy，签发者（Issuer）是 blog-service，有效时间（Expire）为 7200 秒，
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	global.JWTSetting.Expire *= time.Second

	//配置email配置项的读取和映射
	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	return nil
}

//初始化自定义日志
func setupLogger() error {

	global.Logger = logger.NewLogger(&lumberjack.Logger{
		//组合成需要传入的log文件的位置
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2) //#log.LstdFlags:标准logger的初始值

	return nil
}

//初始化完成一个全局变量，不需要返回，值返回error
//初始化mysql链接
func setipDBEngnie() error {
	var err error
	//创建db实例，并将其赋值给项目的全局变量DBEngine
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

//初始化Tracer全局变量，用于链路追踪
func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer(
		"blog-service",
		"127.0.0.1:6831",
	)
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}

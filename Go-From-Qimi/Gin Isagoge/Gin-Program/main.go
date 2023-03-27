package main

import (
	"Gin-Program/dao"
	"Gin-Program/models"
	"Gin-Program/routers"
)

func main() {
	//创建数据库
	//sql:创建表
	//连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	r := routers.SetupRouter()
	r.Run()
}

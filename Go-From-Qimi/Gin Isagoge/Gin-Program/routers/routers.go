package routers

import (
	"Gin-Program/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//告诉gin框架模板文件应用的静态文件去哪里找
	r.Static("/static", "static")
	//告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)
	//v1
	v1Group := r.Group("v1")
	{
		//待办事项
		//添加
		v1Group.POST("todo", controller.CreatTodo)
		//查看:所有和个体
		v1Group.GET("todo", controller.GetTodoList)
		v1Group.GET("todo/:id", func(c *gin.Context) {

		})
		//修改
		v1Group.PUT("todo/:id", controller.UpdateATode)
		//删除
		v1Group.DELETE("todo/:id", controller.DeleteATodo)

	}
	return r
}

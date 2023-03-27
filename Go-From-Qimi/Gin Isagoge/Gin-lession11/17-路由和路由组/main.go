package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	//访问/index 的GET请求会走这一条处理逻辑
	//路由
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})
	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})
	r.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})
	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})

	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			c.JSON(http.StatusOK, gin.H{
				"message": "GET",
			})
		case "POST":
			c.JSON(http.StatusOK, gin.H{"message": "POST"})
		case "PUT":
			c.JSON(http.StatusOK, gin.H{"message": "PUT"})
		case "DELETE":
			c.JSON(http.StatusOK, gin.H{"message": "DELETE"})
		}
	})
	//处理无效URL
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "not found!"})
	})

	//视频的昼夜和详情页
	//路由组概念
	//把公用的前缀提取出来，创建一个路由器
	videGroup := r.Group("/test")
	videGroup.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "/test/index"})
	})
	videGroup.GET("/xxx", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "/test/xxx"})
	})
	videGroup.GET("/00", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "/test/00"})
	})

	r.Run()
}

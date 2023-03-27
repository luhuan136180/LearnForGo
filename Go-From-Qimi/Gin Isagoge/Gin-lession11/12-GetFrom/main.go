package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("12-GetFrom/login.html", "12-GetFrom/index.html")

	//一个请求对应一个响应
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.POST("/login", func(c *gin.Context) {
		//username := c.PostForm("username")
		password, ok := c.GetPostForm("passwd")
		if !ok {
			password = "*****"
		}
		username := c.DefaultPostForm("username", "somebody")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Password": password,
		})
	})
	r.Run(":9090")
}

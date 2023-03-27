package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//重定向
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {

		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")

	})

	r.GET("/a", func(c *gin.Context) {
		c.Request.URL.Path = "/b"
		r.HandleContext(c)
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "b",
		})
	})

	r.Run()
}

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func hello(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello"})
}

func main() {
	r := gin.Default()

	r.GET("/hello", hello)

	r.POST("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "post"})
	})

	r.PUT("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "put",
		})
	})

	r.DELETE("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			",eaasge": "delete",
		})
	})

	r.Run(":9090")
}

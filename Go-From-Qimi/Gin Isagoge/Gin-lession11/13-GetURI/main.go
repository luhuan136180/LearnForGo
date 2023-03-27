package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//path参数返回的都是字符串类型
func main() {
	r := gin.Default()

	r.GET("/:username/:age", func(c *gin.Context) {
		username := c.Param("username")
		password := c.Param("password")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"NAME":   username,
			"PASSWD": password,
			"age":    age,
		})
	})

	r.GET("/blog/:year/:month", func(c *gin.Context) {
		year := c.Param("year")
		month := c.Param("month")
		c.JSON(http.StatusOK, gin.H{
			"Month": month,
			"Year":  year,
		})
	})

	r.Run(":9090")
}

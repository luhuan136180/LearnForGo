package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("14-GetStruct/login.html")

	//http://localhost:8080/user?username=mhx&password=1234
	r.GET("/user", func(c *gin.Context) {
		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})

		} else {
			fmt.Println(u)
			c.JSON(http.StatusOK, gin.H{
				"status":   "ok",
				"UserInfo": u,
			})
		}
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.POST("/form", func(c *gin.Context) {
		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})

		} else {
			fmt.Println(u)
			c.JSON(http.StatusOK, gin.H{
				"status":   "ok",
				"UserInfo": u,
			})
		}
	})

	r.POST("/json", func(c *gin.Context) {
		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})

		} else {
			fmt.Println(u)
			c.JSON(http.StatusOK, gin.H{
				"status":   "ok",
				"UserInfo": u,
			})
		}
	})

	r.Run(":8080")
}

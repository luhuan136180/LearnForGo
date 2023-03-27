package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default() //启动默认路由
	r.GET("/json", func(c *gin.Context) {
		//data := map[string]interface{}{
		//	"name":    "小王子",
		//	"message": "hello world",
		//	"age":     18,
		//}
		data := gin.H{
			"name":    "小王子",
			"message": "hello world",
			"age":     18,
		}
		c.JSON(http.StatusOK, data)
	})

	r.GET("/moreJson", func(c *gin.Context) {
		var msg struct {
			//初始化结构体时，需要大写成public，不能字段小写
			Name    string `json:"user"`
			Message string
			Age     int
		}
		msg.Age = 18
		msg.Name = "mhx"
		msg.Message = "hello 2!"
		c.JSON(http.StatusOK, msg)
	})
	r.Run("9090")
	//fmt.Println("hahahhaha")
}

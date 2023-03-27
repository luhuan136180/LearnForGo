package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		//获取	quertstring 参数
		//1
		//name := c.Query("query") //Query方法：获取请求中的querystring参数
		//2.
		//name := c.DefaultQuery("query", "没有参数")
		//DefaultQuery:当找不到参数名时返回设定的默认值

		//3.func (c *Context) GetQuery(key string) (string, bool)
		name, err := c.GetQuery("query")
		if !err {
			name = "没有该参数"
		}
		age := c.Query("age")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"ok":   "OK",
			"age":  age,
		})
	})

	r.Run(":9090")
}

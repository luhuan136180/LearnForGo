package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func indexHandler(c *gin.Context) {
	fmt.Println("/index")
	//get 一系的方法：get() MustGet()  GetString() GetBool() ....
	name, ok := c.Get("name") //从中间件取参数
	if !ok {
		name = "默认值"
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "index",
		"name": name,
	})
}

//定义一个中间件
func m1(c *gin.Context) {
	fmt.Println("m1 in...")
	start := time.Now()
	c.Set("name", "mhx")
	c.Next()
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("m1 out...")
}

func m2(c *gin.Context) {
	fmt.Println("m2 in...")
	c.Next()
	fmt.Println("m2 out...")
}

func outMiddleware(doCheck bool) gin.HandlerFunc {
	//链接数据库
	//或者一些其他的工作
	return func(c *gin.Context) {
		if doCheck {
			fmt.Println("运行")
			//存放具体的逻辑
			//是否登录的判断
			//if 是登录用户
			//c.Next()
			//else
			//c.Abort()
		} else {
			c.Next()
		}
	}
}
func main() {
	r := gin.Default()

	r.Use(m1, m2, outMiddleware(true))

	r.GET("/index", indexHandler)
	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "shop"})
	})
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "user"})
	})
	//全局中间件任然在局部中生效，局部的中间件只在设定的范围内生效
	//路由组注册局部中间件方法1
	xxGroup := r.Group("/xxx", outMiddleware(true))
	{
		xxGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "hello"})
		})
		xxGroup.GET("/shop", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "shop"})
		})
	}
	//路由组注册局部中间件方法2
	xxGroup2 := r.Group("/ooo")
	xxGroup2.Use(outMiddleware(true))
	{
		xxGroup2.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "hello"})
		})
		xxGroup2.GET("/shop", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "shop"})
		})
	}
	r.Run()
}

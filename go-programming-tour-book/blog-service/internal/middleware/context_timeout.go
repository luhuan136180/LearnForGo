package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

/*
编写一个上下文超时时间控制的中间件来实现这个需求
*/

func ContextTimeout(t time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		//设置当前context的超时时间，并重新赋予给gin。context，
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
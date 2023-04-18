package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
)

//jwt解析中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)
		// 从URL参数或Header中获取token
		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.GetHeader("token")
		}
		//// 如果token为空则为无效参数错误
		if token == "" {
			ecode = errcode.InvalidParams
		} else {
			// 解析token
			_, err := app.ParseToken(token)
			if err != nil {
				// 根据不同的错误类型进行不同的处理
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			}
		}

		// 如果有错误信息，则返回错误响应
		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}

		// 没有错误则继续处理下一个请求
		c.Next()
	}
}

package middleWares

import (
	"Forumv2/controller"
	"Forumv2/pkg/jwt"
	"github.com/gin-gonic/gin"
	"strings"
)

// JWTAuthMiddleware 基于JWT的认证中间件
//将认证操作全部封装进中间件，获取请求的请求体中传入的token，将其解析，并将其中的关键信息存入context上下文中，方便使用和查看
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		//
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" { //及没有login获取token
			controller.ResponseError(c, controller.CodeNeedLogin)
			c.Abort()
			return
		}

		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}

		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1]) //解析
		if err != nil {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}

		// 将当前请求的username信息保存到请求的上下文c上*********
		c.Set(controller.ContextUserAddressKey, mc.UserAddress)

		// 后续的处理函数可以用过c.Get("userID")来获取当前请求的用户信息
		c.Next()
	}
}

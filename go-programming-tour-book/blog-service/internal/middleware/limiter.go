package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
	"github.com/go-programming-tour-book/blog-service/pkg/limiter"
)

//令牌桶限流中间件
func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			//如果当前请求的关键词对应的 Bucket 实例存在
			count := bucket.TakeAvailable(1) // 从 Bucket 中获取一个可用的令牌
			if count == 0 {                  // 如果 Bucket 中没有可用令牌
				response := app.NewResponse(c) // 新建一个 app.Response 实例，用于封装生成服务器响应的方法
				//// 调用 response 实例的 ToErrorResponse 方法，生成一个 "TooManyRequests" 的错误响应，并返回给客户端
				response.ToErrorResponse(errcode.TooManyRequests)

				// // 终止 gin.Context 实例的执行，
				//直接返回响应给客户端。c.Abort() 方法内置实现了 c.AbortWithStatus 方法和 c.AbortWithStatusJSON 方法，其作用是直接返回 HTTP 状态码和 JSON 数据格式的响应。
				c.Abort()
				return
			}
		}

		c.Next() // 继续执行之后的中间件或处理函数
	}
}

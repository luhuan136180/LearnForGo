package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/logger"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

//通过该函数实现了双写，通过body取到值
func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil { //没有报错的时候，将数据也写入body中
		return n, err
	}

	//正常写入
	return w.ResponseWriter.Write(p)
}

//访问日志的中间件
//初始化AccesslogWriter,将其赋予给当前的Writer写入流，并通过制定方法得到我们所需的日志属性，最终写入到我们的日志中，
func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		//使用c.write方法可以往HTTP response中写入数据。
		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter

		//记录请求的开始时间
		beginTime := time.Now().Unix()
		c.Next()
		//记录请求的结束时间，两者记录了请求的耗时。
		endTime := time.Now().Unix()

		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}
		global.Logger.WithFields(fields).Infof("access log: method: %s, status_code: %d, begin_time: %d, end_time: %d",
			c.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime,
		)

	}
}

//自定义异常捕获处理Recovery
//func Recovery() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		defer func() {
//			if err := recover(); err != nil {
//				global.Logger.WithCallersFrames().Errorf("panic recover err: %v", err)
//				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
//				c.Abort()
//			}
//		}()
//		c.Next()
//	}
//
//}

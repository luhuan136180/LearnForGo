package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
)

//该中间件，用于实现将gin和tracer衔接起来，让每次接口的调用都能精确地上报到追踪系统中，
func Tracing() func(c *gin.Context) {
	return func(c *gin.Context) {
		var newCTX context.Context
		var span opentracing.Span
		spanCTX, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(c.Request.Header),
		)
		if err != nil {
			span, newCTX = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				global.Tracer,
				c.Request.URL.Path,
			)
		} else {
			span, newCTX = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				global.Tracer,
				c.Request.URL.Path,
				opentracing.ChildOf(spanCTX),
				opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
			)
		}

		defer span.Finish()

		//获取链路的spanID和traceID信息
		var traceID string
		var spanID string
		var spanContext = span.Context()
		//对spanContext 的jaeger.SpanContext 做断言
		switch spanContext.(type) {
		case jaeger.SpanContext:
			jaegerContext := spanContext.(jaeger.SpanContext)
			traceID = jaegerContext.TraceID().String()
			spanID = jaegerContext.SpanID().String()
		}
		c.Set("X-Trace-ID", traceID)
		c.Set("X-Span-ID", spanID)

		c.Request = c.Request.WithContext(newCTX)
		c.Next()
	}
}

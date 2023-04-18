package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/middleware"
	"github.com/go-programming-tour-book/blog-service/internal/routers/api"
	v1 "github.com/go-programming-tour-book/blog-service/internal/routers/api/v1"
	"github.com/go-programming-tour-book/blog-service/pkg/limiter"
	"net/http"
	"time"

	//一下是接口文档相关依赖
	_ "github.com/go-programming-tour-book/blog-service/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	Key:          "/auth",
	FillInterval: time.Second,
	Capacity:     10,
	Quantum:      10,
})

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}

	//限流中间件
	r.Use(middleware.RateLimiter(methodLimiters))

	//统一超时时间中间件**注：此处我们将其写死为60秒
	r.Use(middleware.ContextTimeout(60 * time.Second))

	//注册中间件：翻译器
	r.Use(middleware.Translations())
	r.Use(middleware.Tracing()) //注册中间件：链路跟踪
	//接口文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	tag := v1.NewTag()
	article := v1.NewArticle()
	upload := api.NewUpload()

	r.POST("/upload/file", upload.UploadFile) //上传文件对应路由
	//以下：当请求 URL 中匹配到 /static 的时候，就会自动到全局配置中的静态文件所在目录
	//global.AppSetting.UploadSavePath 下寻找文件。如果找到该文件，Gin 会自动将该文件发送给请求方作为响应。
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath)) //Gin 框架中，前端的静态文件可以直接访问到，并且可以在静态文件中包含图片、CSS 样式表、JavaScript 文件等等。

	//获取 Token
	r.POST("/auth", api.GetAuth)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT()) //对该组实行中间件约束

	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update) //备注：因为state是uint8类型，而且没有设计专用的绑定参数函数，所以使用有问题
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}

	return r
}

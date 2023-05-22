package routers

import (
	"Forumv1/controller"
	"Forumv1/logger"
	"Forumv1/middleWares"
	"github.com/gin-gonic/gin"
)

func SetUpRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) //设置成发布模式
	}

	r := gin.New()
	//使用自己编写的logger
	//r.Use(logger.GinLogger(), logger.GinRecovery(true), middleWares.RateLimitMiddleware(2*time.Second, 1))
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//
	v1 := r.Group("/api/v1")
	v1.GET("/helloworld", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	//注册
	v1.POST("/signup", controller.SignUpHandler)
	//登录
	v1.POST("/login", controller.LoginHandler)

	//继续访问该项目的内容需要先登录，所以需要认证处理，添加中间件
	v1.Use(middleWares.JWTAuthMiddleware()) //应用JWT认证中间件
	{
		//testAPI
		v1.GET("/Hello", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "hello,user!",
			})
		})
		//创建主题社区
		v1.POST("/topic/create", controller.CreateTopicHandler)
		//查询社区的内容（所有社区的简述）
		v1.GET("/topics", controller.GetTopicListHandler)
		//查询单个社区的帖子
		v1.GET("/topic/id/:id", controller.TopicDetailBYIDHandler)
		v1.GET("/topic/name/:name", controller.TopicDetailBYNameHandler)
		//创建帖子
		v1.POST("/post/create", controller.CreatePostHandler)
		//删除帖子

		//查询单个帖子的主贴
		v1.GET("post/bykey/:key", controller.GetPostByKeyHandler)
		//根据主题，查找帖子
		v1.GET("post/bytopicid/:topicid/", controller.GetPostsByTopicIDHandler)
		//查询一些帖子---模糊查询--对内容版本
		v1.GET("post/like-content/:word", controller.GetPostByContentLIKEHandler)

		//对帖子发表评论
		v1.POST("post/:postkey/response/", controller.CreateResponseHandler)

		//查询一个帖子的主贴加回复等等等
		v1.GET("post/getallpost/bypost_key", controller.GetAllPostsByPostKeyHandler)

		//查询一个用户的余额
		v1.GET("/user/balance", controller.GetUserBalanceHandler)
		//进行交易，改变用户余额--添加
		v1.PUT("/user/add_balance", controller.AddBalanceHandler)
		//进行交易，改变用户余额--支出
		v1.PUT("/user/sub_amount", controller.SubBalanceHandler)
		//修改用户名

		//上传文件？？---纯上传API
		v1.POST("/file", controller.UploadFile)
		//上传文件，同时记录上传用户等信息
		v1.POST("/file/user", controller.UploadFileWithAuthor)

		//获取本地存储的文件
	}
	return r
}

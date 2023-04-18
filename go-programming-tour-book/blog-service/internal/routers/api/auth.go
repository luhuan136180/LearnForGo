package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/service"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
)

//校验及获取入参后，绑定并获取到的 app_key 和 app_secrect 进行数据库查询，检查认证信息是否存在，若存在则进行 Token 的生成并返回。
func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	//fmt.Fprintf(os.Stdout, "参数：", param.AppKey, param.AppSecret)
	svc := service.New(c.Request.Context())
	//fmt.Println(param.AppKey, param.AppSecret)
	err := svc.CheckAuth(&param)

	if err != nil {
		global.Logger.Errorf("svc.CheckAuth err : %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
	}
	//传入的appkey和appscret是合法的，于是使用这两者来进行token生成
	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf("app.GenerateToken err:!! %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})
}

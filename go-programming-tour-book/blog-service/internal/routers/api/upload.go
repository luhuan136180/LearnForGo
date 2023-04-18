package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/service"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/go-programming-tour-book/blog-service/pkg/convert"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
	"github.com/go-programming-tour-book/blog-service/pkg/upload"
)

type Upload struct {
}

func NewUpload() Upload {
	return Upload{}
}

//通过 c.Request.FormFile 读取入参 file 字段的上传文件信息，
//并利用入参 type 字段作为所上传文件类型的确立依据（也可以通过解析上传文件后缀来确定文件类型），
//最后通过入参检查后进行 Serivce 的调用，完成上传和文件保存，返回文件的展示地址。
func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c) //封装传入的context上下文
	//用于获取上传文件参数中的某个文件，参数是上传文件的字段名。返回值是一个 multipart.File 类型的文件对象和一个 multipart.FileHeader 类型的文件头信息对象。可以用于获取用户上传的文件。
	file, fileHeader, err := c.Request.FormFile("file") //**在实际应用中一般不建议使用 c.Request.FormFile。使用 c.FormFile 可以避免对整个请求体进行解析，也更符合业界的最佳实践。
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	//c.Request.Context() 返回的是 HTTP 请求上下文的 context.Context 对象，它所负责的是控制请求的上下文，例如：超时控制、 deadline、密钥等等。
	svc := service.New(c.Request.Context()) //C对象和 Request.Context() 是从不同的角度来看待请求上下文的，并且它们使用的场景也不同。换句话说，C对象更多的是可以让您在处理请求时直接操作请求的信息，而 Request.Context() 会更多地用于构建请求生命周期内的跨链接信息和状态。
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf("svc.UploadFile err:%v", err)                             //记入日志
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error())) //返回相应（错误）
		return
	}

	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}

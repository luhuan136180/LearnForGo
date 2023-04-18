package service

import (
	"errors"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/upload"
	"mime/multipart"
	"os"
)

//用于将上传文件的工具库使用到具体的业务接口中的函数封装
type FileInfo struct {
	Name      string
	AccessUrl string
}

//上传文件的函数，主要作用是将前端浏览器传来的文件保存到服务器上，并返回该文件在服务器上的访问路径。
//该函数的参数包括文件类型、文件、文件头，其中文件类型是一个枚举类型，用于指明文件的类型（例如图片、音频、视频等），
//文件是前端浏览器传来的文件，文件头则包含了该文件的信息。函数内部会先检查文件后缀名是否在支持列表中，再检查文件是否超过最大限制大小，
//随后会获取将要保存的文件路径并检查路径是否存在和权限是否足够，最后保存文件并返回该文件在服务器上的访问路径。
func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename)
	if !upload.CheckContainExt(fileType, fileName) { //检查文件后缀名是否在支持的列表中
		//检查是否是项目认可的后缀格式
		return nil, errors.New("file suffix is not supported.")
	}
	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit.")
	}

	//
	uploadSavePath := upload.GetSavePath()    //获取保存文件的路径；
	if upload.CheckSavePath(uploadSavePath) { //检查保存文件的路径是否存在；
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory.")
		}
	}
	if upload.CheckPermission(uploadSavePath) { //查看文件权限是否足够
		return nil, errors.New("insufficient file permissions.")
	}

	dst := uploadSavePath + "/" + fileName
	if err := upload.SaveFile(fileHeader, dst); err != nil { //保存文件
		return nil, err
	}

	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}

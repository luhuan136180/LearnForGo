package logic

import (
	"Forumv1/dao/mysql"
	"Forumv1/models"
)

func CreateFileWithAuthor(fileinfo *models.FileUpload) (err error) {
	//存储到数据库

	err = mysql.CreateFileWithAuthor(fileinfo)
	return
}

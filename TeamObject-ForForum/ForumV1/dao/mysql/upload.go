package mysql

import "Forumv1/models"

func CreateFileWithAuthor(fileinfo *models.FileUpload) (err error) {
	sqlStr := `insert into file_info(
	filename, file_id, user_id)
	values (?,?,?)
	`
	_, err = Db.Exec(sqlStr, fileinfo.FileName, fileinfo.FileID, fileinfo.UserId)
	return
}

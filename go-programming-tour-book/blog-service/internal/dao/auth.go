package dao

import (
	"github.com/go-programming-tour-book/blog-service/internal/model"
)

//获取数据库中存储的本服务器认证密钥，认证key等信息
func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{AppKey: appKey, AppSecret: appSecret}
	return auth.Get(d.engine)
}

package models

//登录参数
type ParamLogin struct {
	Username string `json:"username"binding:"required"`
	Password string `json:"password"binding:"required"`
}

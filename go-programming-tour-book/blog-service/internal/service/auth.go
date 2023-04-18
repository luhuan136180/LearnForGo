package service

import "errors"

//用于接口入参校验，参数必填
type AuthRequest struct {
	AppKey    string `form:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" binding:"required"`
}

//校验：使用客户端传入的认证信息作为帅选条件在数据库中获取数据行，以此根据是否取到认证信息来判断
// CheckAuth 函数是用来校验应用程序的 AppKey 和 AppSecret 的合法性的
func (svc *Service) CheckAuth(param *AuthRequest) error {
	auth, err := svc.dao.GetAuth(param.AppKey, param.AppSecret)
	if err != nil {
		// 如果在从数据库中获取 Auth 记录时出错，则返回错误信息 err
		return err
	}
	if auth.ID > 0 {
		// 如果从数据库中获取到的 Auth 记录的 ID 大于 0，说明该 Auth 记录存在，返回 nil 标识校验通过
		return nil
	}
	// 如果 ID <= 0，说明没有找到对应的 Auth 记录，返回错误信息 "auth info does not exist."
	return errors.New("auth info does not exist.")
}

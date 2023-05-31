package logic

import (
	"Forumv1/dao/mysql"
	"Forumv1/models"
	"Forumv1/pkg/jwt"
	"Forumv1/pkg/snowflake"
	"fmt"
)

func Login(p *models.ParamLogin) (token string, err error) {
	//对数据进行校验-1.是否存在
	user := &models.User{
		Name:     p.Username,
		Password: p.Password,
	}
	exit, err := mysql.CheckUserExist(user.Name)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	//判断该用户是否已经存储
	if !exit { //不存在
		//表中没有数据，进入注册流程
		err = SignUp(user) //现在存在了
		if err != nil {
			return "", err
		}
		return jwt.GenToken(user.UserID, user.Name)
	}
	//存在，正常进入登录，传递的是指针,
	if err := mysql.Login(user); err != nil {
		return "", err
	}

	return jwt.GenToken(user.UserID, user.Name)
}

func SignUp(user *models.User) (err error) {
	userID := snowflake.GenID()
	user.UserID = userID
	fmt.Println(user)
	//保存进数据库
	return mysql.InsertUser(user)
}

func GetUserBalance(userid int64) (data *models.Balance, err error) {
	return mysql.GetUserBalance(userid)
}

func AddBalance(transaction *models.AmountChange) (data *models.AmountChange, err error) {
	return mysql.AddBalance(transaction)
}

func SubBalance(transaction *models.AmountChange) (data *models.AmountChange, err error) {
	return mysql.SubBalance(transaction)
}

package controller

import (
	"Forumv1/dao/mysql"
	"Forumv1/logic"
	"Forumv1/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	//1.获取参数

	//2.业务处理

	//3.返回相应
	//c.JSON(200, gin.H{"data": "success"})
	ResponseSuccess(c, "登录成功")
}

//从前端获取登录参数参数————参数应该为：用户名：公钥,没有密码???
//本次版本仍写上一个密码
func LoginHandler(c *gin.Context) {
	//fmt.Println("0")
	//获取参数，进行校验
	p := new(models.ParamLogin) //创建实例
	if err := c.ShouldBind(p); err != nil {
		//
		//fmt.Println("1")
		zap.L().Error("Login with invalid param", zap.Error(err))
		tanser, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(tanser.Translate(trans)))
		return
	}
	//业务处理
	token, err := logic.Login(p)
	//fmt.Println(token)
	if err != nil {
		zap.L().Error("Logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		} else if errors.Is(err, mysql.ErrorInvalidPassword) {
			ResponseError(c, CodeInvalidPassword)
			return
		}
		ResponseErrorWithMsg(c, CodeMysql, err.Error())
		return
	}

	//返回请求
	ResponseSuccess(c, token)
}

//查询用户的余额
func GetUserBalanceHandler(c *gin.Context) {
	userID, err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	//fmt.Println(userID)
	//通过用户id查询用户余额
	data, err := logic.GetUserBalance(userID)
	if err != nil {
		zap.L().Error("Logic.GetUserBalance failed", zap.String("userid", string(data.UserID)), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		} else if errors.Is(err, mysql.ErrorInvalidPassword) {
			ResponseError(c, CodeInvalidPassword)
			return
		}
		ResponseErrorWithMsg(c, CodeMysql, err.Error())
		return
	}
	fmt.Println(data)
	ResponseSuccess(c, data)
}

//交易--收入--只能自己给自己充钱
func AddBalanceHandler(c *gin.Context) {
	//使用json绑定
	transaction := new(models.AmountChange)
	if err := c.ShouldBindJSON(transaction); err != nil {
		zap.L().Error("ADDBalance is Err", zap.Error(err))
		tanser, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(tanser.Translate(trans)))
		return
	}

	if transaction.Amount <= 0 {
		ResponseErrorWithMsg(c, CodeInvalidParam, "Amount必须是整数")
		return
	}

	//userID, err := GetCurrentUser(c)
	//if err != nil {
	//	zap.L().Error("GetCurrentUser failed", zap.Error(err))
	//	ResponseError(c, CodeNeedLogin)
	//	return
	//}

	//进入业务流程
	data, err := logic.AddBalance(transaction)
	if err != nil {
		zap.L().Error("logic.AddBalance(transaction) is failed", zap.Error(err))
		ResponseError(c, CodeMysql)
		return
	}

	//成功响应
	ResponseSuccess(c, data)
}

//支出
func SubBalanceHandler(c *gin.Context) {
	//
	transaction := new(models.AmountChange)
	if err := c.ShouldBindJSON(transaction); err != nil {
		zap.L().Error("SubBalance is Err", zap.Error(err))
		tanser, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(tanser.Translate(trans)))
		return
	}
	fmt.Println(transaction)
	if transaction.Amount >= 0 {
		ResponseErrorWithMsg(c, CodeInvalidParam, "Amount必须是负数")
		return
	}

	data, err := logic.SubBalance(transaction)
	if err != nil {
		zap.L().Error("logic.SubBalance(transaction) is failed", zap.Error(err))
		ResponseError(c, CodeMysql)
		return
	}
	//成功响应
	ResponseSuccess(c, data)
}

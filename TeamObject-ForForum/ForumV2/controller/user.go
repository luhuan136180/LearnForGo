package controller

import (
	"Forumv2/dao/mysql"
	"Forumv2/logic"
	"Forumv2/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strconv"
)

//登录
//获取参数为一个签名和一个公钥地址
func LoginHandler(c *gin.Context) {

	//获取参数
	userLogin := new(models.Login)
	//
	//
	if err := c.ShouldBindJSON(userLogin); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		tanser, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(tanser.Translate(trans)))
		return
	}
	//
	fmt.Println("userLogin", userLogin)
	user := new(models.User)
	user.UserAddress = userLogin.UserAddress
	fmt.Println("user", user)
	//业务处理
	token, err := logic.Login(user)
	//fmt.Println(token)
	if err != nil {
		zap.L().Error("Logic.Login failed", zap.String("username", user.UserAddress), zap.Error(err))
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

func GetUserBalanceHandler(c *gin.Context) {
	user_address := c.Param("user_address")
	data, err := logic.GetUserBalance(user_address)
	if err != nil {
		zap.L().Error("Logic.GetUserBalance failed", zap.String("user_address", user_address), zap.Error(err))
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
	ResponseSuccess(c, data)
}

func AddBalanceHandler(c *gin.Context) {
	user_address := c.Param("user_address")
	amountS := c.Query("amount")
	amount, _ := strconv.Atoi(amountS)
	//
	data, err := logic.AddUserBalance(user_address, amount)
	if err != nil {
		zap.L().Error("logic.AddUserBalance(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func SubBalanceHandler(c *gin.Context) {
	user_address := c.Param("user_address")
	amountS := c.Query("amount")
	amount, _ := strconv.Atoi(amountS)

	data, err := logic.SubUserBalance(user_address, amount)
	if err != nil {
		zap.L().Error("logic.AddUserBalance(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

//
func GetUserInformation(c *gin.Context) {
	user_address := c.Param("user_address")
	data, err := logic.GetUserInformation(user_address)
	if err != nil {
		zap.L().Error("logic.AddUserBalance(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

//获取用户的所有拥有的皮肤
func GetAllSkinByUserHandler(c *gin.Context) {
	user_address := c.Param("user_address")
	data, err := logic.GetAllSkinByUser(user_address)
	if err != nil {
		zap.L().Error("logic.GetAllSkinByUser(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

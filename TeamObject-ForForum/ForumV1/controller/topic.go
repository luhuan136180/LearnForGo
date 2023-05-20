package controller

import (
	"Forumv1/logic"
	"Forumv1/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

//跟主题相关

//创建一个主题
func CreateTopicHandler(c *gin.Context) {
	//创建主题，需要输入主题名称，主题的内容概述...待添加
	//获取参数, 先假定以json格式传入,
	topic := new(models.TopicDetail)                //创建
	if err := c.ShouldBindJSON(topic); err != nil { //绑定数据
		zap.L().Debug("**c.shouldBindJson(p) error", zap.Any("err", err))
		zap.L().Error("***create post with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	//业务流程---将topic结构体传入，验证该主题是否已经存在等，然后向数据库注入
	if err := logic.CreateTopic(topic); err != nil {
		zap.L().Error("logic.CreatTopic(topic) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	//响应
	ResponseSuccess(c, "添加成功")
}

//查询所有主题
func GetTopicListHandler(c *gin.Context) {
	//查询到所有主题（id，name）以列表形式返回
	data, err := logic.GetTopicList()
	if err != nil {
		zap.L().Error("logic.GetTopicList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) //不轻易将数据库层的东西暴露给服务
		return
	}
	ResponseSuccess(c, data)
}

//根据id查询单个主题的详细信息
func TopicDetailBYIDHandler(c *gin.Context) {
	idStr := c.Param("id")                     //获取url中的参数
	id, err := strconv.ParseInt(idStr, 10, 64) //string转int64
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	//根据获取的id获取社区详情
	data, err := logic.GetTopicDetailByID(id)
	if err != nil {
		zap.L().Error("logic.TopicDetailBYIDHandler() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, data)
}

func TopicDetailBYNameHandler(c *gin.Context) {
	name := c.Param("name")
	data, err := logic.GetTopicDetailByName(name)
	if err != nil {
		zap.L().Error("logic.GetTopicDetailByName() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, data)
}

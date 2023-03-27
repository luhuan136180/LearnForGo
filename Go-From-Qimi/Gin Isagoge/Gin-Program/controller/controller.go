package controller

import (
	"Gin-Program/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	 url     --> controller  --> logic   -->    model
	请求来了  -->  控制器      --> 业务逻辑  --> 模型层的增删改查
*/

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreatTodo(c *gin.Context) {
	//前端页面填写待办事项
	//1.从请求中将数据拿出来
	var todo models.Todo
	c.BindJSON(&todo)
	//2.存入数据库
	if err := models.CreateAtodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": todo,
		})
	}
	//3.反应响应
}

func GetTodoList(c *gin.Context) {
	if todoList, err := models.GetAllTodo(); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateATode(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"Error": "无效的ID，重新输入"})
	}
	//查找
	todo, err := models.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	//获取从网页传来的更新数据
	c.BindJSON(todo)
	//将实例出入数据库中完成更新
	if err = models.UpdateATodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}

}

func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效id"})
	}
	if err := models.DeleteATodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "delete"})
	}
}

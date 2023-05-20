package logic

import (
	"Forumv1/dao/mysql"
	"Forumv1/models"
	"Forumv1/pkg/snowflake"
	"fmt"
)

//创建主题
func CreateTopic(topic *models.TopicDetail) (err error) {
	//检验该主题，在数据库中查询，该主题是否已经注册
	err = mysql.CheckTopicExist(topic.Name)
	if err != nil { //检查过程中出现错误，退出
		fmt.Println(err)
		return err
	}
	//不存在，可以放心添加
	topic.ID = snowflake.GenID()

	//进入添加数据流程
	err = mysql.CreateTopic(topic)
	if err != nil {
		return err
	}
	return nil
}

func GetTopicList() ([]*models.Topic, error) {
	//查询数据
	return mysql.GetTopicList()
}

func GetTopicDetailByID(id int64) (*models.TopicDetail, error) {
	return mysql.GetTopicDetailByID(id)
}

func GetTopicDetailByName(name string) (*models.TopicDetail, error) {
	return mysql.GetTopicDetailByName(name)
}

package mysql

import (
	"Forumv1/models"
	"database/sql"
	"go.uber.org/zap"
)

func CheckTopicExist(topicName string) (err error) {
	sqlStr := "select count(topic_id) from topic where topic_name=?"
	var count int
	if err := Db.Get(&count, sqlStr, topicName); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return nil
}

func CreateTopic(topic *models.TopicDetail) (err error) {
	sqlStr := `insert into topic(
	topic_id, topic_name, introduction)
	values (?,?,?)
	`
	_, err = Db.Exec(sqlStr, topic.ID, topic.Name, topic.Introduction)
	return err
}

//查询所有的topic
func GetTopicList() (topicList []*models.Topic, err error) {
	sqlStr := "select topic_id,topic_name from topic"
	if err = Db.Select(&topicList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}

//查询单个topic
func GetTopicDetailByID(id int64) (topic *models.TopicDetail, err error) {
	//初始化
	topic = new(models.TopicDetail)
	sqlStr := "select" +
		" topic_id,topic_name,introduction,create_time" +
		" from topic" +
		" where id=?"
	if err = Db.Get(topic, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return
}

func GetTopicDetailByName(name string) (topic *models.TopicDetail, err error) {
	topic = new(models.TopicDetail)
	sqlStr := "select" +
		" topic_id,topic_name,introduction,create_time" +
		" from topic" +
		" where topic_name=?"
	if err = Db.Get(topic, sqlStr, name); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return
}

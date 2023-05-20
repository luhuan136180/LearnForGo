package mysql

import (
	"Forumv1/models"
	"fmt"
	"testing"
)

func TestCheckTopicExist(t *testing.T) {
	topic := &models.TopicDetail{
		Name:         "go",
		Introduction: "testing",
	}
	err := CheckTopicExist(topic.Name)
	if err != nil {
		fmt.Println("err:", err)
	}
}

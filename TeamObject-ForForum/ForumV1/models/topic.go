package models

import "time"

type Topic struct {
	ID   int64  `json:"topic_id,string"db:"topic_id"`
	Name string `json:"name"db:"topic_name"`
}

type TopicDetail struct {
	ID   int64  `json:"topic_id,string"db:"topic_id"`
	Name string `json:"name"binding:"required"db:"topic_name"`
	//
	Introduction string `json:"introduction,omitempty"binding:"required"db:"introduction"`
	//创建时间
	CreateTime time.Time `json:"create_time"db:"create_time"`
}

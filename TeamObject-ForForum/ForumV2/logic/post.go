package logic

import (
	"Forumv2/dao/mysql"
	"Forumv2/models"
	"Forumv2/pkg/snowflake"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"strconv"

	"time"
)

func CreatePost(post *models.CtreatePost) (data *models.CreateResponse, err error) {

	post.PostID = snowflake.GenID()
	data = new(models.CreateResponse)
	//生成hash值
	value := post.Content + post.AuthorAddress + time.Now().String()
	hash := encryptContent(value)
	post.PostKey = hash //内容的哈希值
	post.TopicID = 1    //默认值，此处固定
	post.Status = 1
	//数据处理完成，保存进数据库,
	//fmt.Println(*post)
	err = mysql.CreatePost(post)
	if err != nil {
		return
	}
	data.Title = post.Title
	data.PostKey = hash
	return
}

func encryptContent(value string) string {
	data := []byte(value)
	hash := sha512.Sum512(data)
	return hex.EncodeToString(hash[:])
}

func GetPostsList(page, size int64) (data []*models.GetPost, err error) {
	posts, err := mysql.GetPostsList(page, size)
	if err != nil {
		return nil, err
	}
	fmt.Println(posts)
	return posts, nil
}

func GetPostByContentLIKE(word string) (data []*models.GetPost, err error) {
	data, err = mysql.GetPostByContentLIKE(word)
	//fmt.Println(data)
	if err != nil {
		return nil, err
	}
	return
}

func GetPostByTitleLIKE(word string) (data []*models.GetPost, err error) {
	data, err = mysql.GetPostByTitleLIKE(word)
	//fmt.Println(data)
	if err != nil {
		return nil, err
	}
	return
}

func GetPostByPostID(postid string) (data []*models.GetPost, err error) {
	postID, _ := strconv.ParseInt(postid, 10, 64)
	data, err = mysql.GetPostByPostID(postID)
	if err != nil {
		return nil, err
	}
	return
}
func CreateResponseByPostID(post *models.CtreatePost) (data *models.CreateResponse, err error) {
	value := post.Content + post.AuthorAddress + time.Now().String()
	hash := encryptContent(value)
	post.PostKey = hash //内容的哈希值
	post.TopicID = 1    //默认值，此处固定
	post.Status = 0     //回复

	data = new(models.CreateResponse)
	if err = mysql.CreateResponseByPostID(post); err != nil {
		return nil, err
	}
	data.PostKey = hash
	return
}

package logic

import (
	"Forumv1/dao/mysql"
	"Forumv1/models"
	"Forumv1/pkg/snowflake"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

const secret = "hash"

func CreatePost(p *models.Post) (err error) {
	//利用雪花算法 生成post id
	p.ID = snowflake.GenID()
	//生成哈希值--当前未使用sha，转而使用MD5加密
	value := string(p.AuthorID) + p.Content + time.Now().String()
	hash := encryptContent(value)
	//
	//fmt.Println("hash=", hash)

	p.PostKey = hash

	//保存到数据库中
	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	return
}

func encryptContent(value string) string {
	h := md5.New()          //
	h.Write([]byte(secret)) //密钥,
	//EncodeToString：返回字符串
	return hex.EncodeToString(h.Sum([]byte(value)))
}

//本版本为正常且修改了上一版本问题的函数，
func encryptContent2(value string) string {
	h := md5.New()         //
	h.Write([]byte(value)) //密钥,
	//EncodeToString：返回字符串
	return hex.EncodeToString(h.Sum([]byte(secret)))
}
func GetPostByKey(key string) (post *models.GetPostByKey, err error) {
	return mysql.GetPostByKey(key)
}

func GetPostsByTopicID(page, size int64, topticid string) (data []*models.GetPostByKey, err error) {
	posts, err := mysql.GetPostsByTopicID(page, size, topticid)
	//fmt.Println(posts)
	if err != nil {
		return nil, err
	}
	data = make([]*models.GetPostByKey, 0, len(posts))
	for _, post := range posts {
		data = append(data, post)
	}
	return
}

func GetPostByContentLIKE(word string) (data []*models.GetPostByKey, err error) {
	posts, err := mysql.GetPostByContentLIKE(word)
	if err != nil {
		return nil, err
	}
	data = make([]*models.GetPostByKey, 0, len(posts))
	for _, post := range posts {
		data = append(data, post)
	}

	return
}

func CreatePostResponse(response *models.ResponseCreate) (err error) {
	//计算出postkey
	value := string(response.AuthorID) + response.Content + time.Now().String()
	hash := encryptContent2(value)
	fmt.Println("hash:", hash)
	response.PostKey = hash
	err = mysql.CreatePostResponse(response)
	if err != nil {
		return err
	}
	return
}

func GetAllPostsByPostKey(postKeyList *models.AutoGenerated) (data []*models.GetPostwithRes, err error) {
	var ListPostKeys []string
	for _, postkey := range *postKeyList {
		ListPostKeys = append(ListPostKeys, postkey.Postkey)
	}
	fmt.Println(ListPostKeys)
	data, err = mysql.GetAllPostsByPostKey(ListPostKeys)
	if err != nil {
		return nil, err
	}
	return
}

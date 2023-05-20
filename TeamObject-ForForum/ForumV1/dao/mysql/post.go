package mysql

import (
	"Forumv1/models"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
)

//创建帖子
func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(
	post_id, title, content, author_id, topic_id,post_key)
	values (?,?,?,?,?,?)
	`
	_, err = Db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.TopicID, p.PostKey)
	return
}

func GetPostByKey(postkey string) (post *models.GetPostByKey, err error) {
	post = new(models.GetPostByKey)
	sqlStr := `select post_id,title,content,author_id,topic_id from post where post_key=?`
	if err = Db.Get(post, sqlStr, postkey); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return
}

func GetPostsByTopicID(page, size int64, topicid string) (posts []*models.GetPostByKey, err error) {
	topicID, _ := strconv.Atoi(topicid)
	sqlStr := `select title,content,author_id,topic_id from post where topic_id=? limit ?,?`
	posts = make([]*models.GetPostByKey, 0, size)
	err = Db.Select(&posts, sqlStr, topicID, (page-1)*size, size)
	//fmt.Println(posts)
	return posts, err
}

func GetPostByContentLIKE(word string) (posts []*models.GetPostByKey, err error) {
	keyword := "%" + word + "%"
	sqlStr := `select title,content,author_id,topic_id from post where content LIKE ?`
	//此处初始化有问题，当模糊查询结果过多时，会反复扩大切片的容量，开销巨大
	posts = make([]*models.GetPostByKey, 0, 4)

	err = Db.Select(&posts, sqlStr, keyword)

	return posts, err
}

func CreatePostResponse(response *models.ResponseCreate) error {
	sqlStr := `insert into post(
	post_id,  content, author_id, topic_id,post_key)
	values (?,?,?,?,?)
	`
	_, err := Db.Exec(sqlStr, response.ID, response.Content, response.AuthorID, response.TopicID, response.PostKey)
	return err
}

func GetAllPostsByPostKey(ListPostKeys []string) (postList []*models.GetPostwithRes, err error) {
	query, args, err := sqlx.In("select post_id,author_id,title,content,topic_id from post where post_key IN (?)", ListPostKeys)
	if err != nil {
		return
	}
	postList = make([]*models.GetPostwithRes, 0, len(ListPostKeys))
	err = Db.Select(&postList, query, args...)
	if err != nil {
		fmt.Println("err 2 here")
	}
	return
}

package mysql

import (
	"Forumv1/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
)

const secret = "forum-v1"

// CheckUserExist 检查指定用户名的用户是否存在
//true :表示存在，  false：表示不存在
func CheckUserExist(username string) (bool bool, err error) {
	sqlStr := "select count(user_id) from user where username=?"
	var count int
	if err := Db.Get(&count, sqlStr, username); err != nil {
		return false, err
	}
	if count > 0 { //表中有数据，存在
		return true, nil
	}
	return false, nil
}

func InsertUser(user *models.User) (err error) {
	//对密码加密
	user.Password = encryptPassword(user.Password)
	user.Balance = 0
	//执行SQL语句入库
	sqlStr := "insert into user(user_id,username,password,balance) values(?,?,?,?)"
	_, err = Db.Exec(sqlStr, user.UserID, user.Name, user.Password, user.Balance)
	return
}

// encryptPassword 密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()          //
	h.Write([]byte(secret)) //密钥,
	//EncodeToString：返回字符串
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *models.User) (err error) {
	oPassword := user.Password //记录原始密码
	sqlStr := "select user_id,username,password from user where username=?"
	err = Db.Get(user, sqlStr, user.Name)
	if err == sql.ErrNoRows {
		//没有查询到
		return ErrorUserNotExist
	}
	if err != nil {
		//查询数据库失败
		return err
	}
	//判断密码是否正确
	password := encryptPassword(oPassword) //对密码加密，与库中比较
	//fmt.Println("opassword=", password)
	//fmt.Println("DBpassword=", user.Password)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	//fmt.Println("success ,相等")
	return
}

func GetUserBalance(userid int64) (data *models.Balance, err error) {
	data = new(models.Balance)
	sqlStr := "select user_id,balance from user where user_id=?"
	err = Db.Get(data, sqlStr, userid)
	if err == sql.ErrNoRows {
		//没有查询到
		return nil, ErrorUserNotExist
	}
	if err != nil {
		//查询数据库失败
		return nil, err
	}

	return
}

func AddBalance(transcation *models.AmountChange) (data *models.AmountChange, err error) {
	data = new(models.AmountChange)
	//fmt.Println(transcation)
	sqlStr := "select user_id,balance from user where user_id=?"
	err = Db.Get(data, sqlStr, transcation.UserID)
	if err != nil {
		return nil, err
	}
	data.Amount += transcation.Amount
	sqlStr2 := "update user set balance=? where user_id=?"
	_, err = Db.Exec(sqlStr2, data.Amount, data.UserID)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func SubBalance(transaction *models.AmountChange) (data *models.AmountChange, err error) {
	data = new(models.AmountChange)
	sqlStr := "update user set balance=balance+? where user_id=?"
	_, err = Db.Exec(sqlStr, data.Amount, data.Amount)
	if err != nil {
		return nil, err
	}
	return data, nil
}

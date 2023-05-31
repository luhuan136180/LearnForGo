package models

type User struct {
	UserID   int64  `db:"user_id,string"`
	Name     string `db:"username"`
	Password string `db:"password"`
	Balance  int    `db:"balance"`
}

//注意：当前没有为用户写初始化余额的函数，目前版本在数据库中手动添加金额
type Balance struct {
	UserID  int64 `json:"user_id"db:"user_id"`
	Balance int   `json:"balance"db:"balance"`
}

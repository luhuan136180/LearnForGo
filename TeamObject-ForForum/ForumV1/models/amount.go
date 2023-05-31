package models

type AmountChange struct {
	UserID int64 `json:"user_id,string"db:"user_id"` //用户id
	Amount int   `json:"amount,string"db:"balance"`  //金额
}

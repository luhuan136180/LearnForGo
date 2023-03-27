package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}
type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/study-gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {

	}
	
	// 自动迁移
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Product{})

	//user := User{Name: "mhx5", Age: 18, Birthday: time.Now()}

	//db.Create(&user)
	//fmt.Println(user.ID)
	//fmt.Println(result.Error)
	//fmt.Println(result.RowsAffected)
	//fmt.Println(user.ID)

	//创建记录并更新给出的字段
	//db.Select("Name", "Age", "CreatedAt").Create(&user)

	//创建一个记录且忽略传递给略去的字段值
	//db.Omit("Name", "Age", "CreatedAt").Create(&user)

	//批量插入
	//var users = []User{{Name: "asd", Age: 12}, {Name: "23asda"}, {Age: 100}}
	//db.Create(&users)
	//for _, user := range users {
	//	fmt.Println(user.ID)
	//}

}

package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserInfo2 struct {
	ID   int64
	Name string `gorm:"default:'毛浩昕'"`
	Age  int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/study-gorm?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//.将表和模型联系起来
	db.AutoMigrate(&UserInfo2{})

	//3.创建
	//	u := UserInfo2{
	//		Name: "",
	//		Age:  29,
	//	}
	//u := UserInfo2{Name: sql.NullString{String: "", Valid: true}, Age: 21}
	//
	//db.Create(&u)
	//fmt.Println(u.ID)

	//4.查询
	var user UserInfo2
	db.First(&user) //查询第一条记录
	// SELECT * FROM users ORDER BY id LIMIT 1;
	fmt.Printf("user:%#v\n", user)
	var user2 UserInfo2
	//db.Take(&user2) // 获取一条记录，没有指定排序字段
	//// SELECT * FROM users LIMIT 1;
	//fmt.Printf("user:%#v\n", user2)

	db.Last(&user2) // 获取最后一条记录（主键降序）
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	fmt.Printf("user:%#v\n", user2)

	var user3 UserInfo2
	// 查询指定的某条记录(仅当主键为整型时可用)
	db.First(&user3, 4)
	//// SELECT * FROM users WHERE id = 10;
	fmt.Printf("user:%#v\n", user3)

	//where 条件查询
	var u UserInfo2
	var u2 UserInfo2
	var u3 UserInfo2
	db.Where("name=?", "mhx").First(&u)
	fmt.Printf("user:%#v\n", u)

	db.Where("name=?", "毛浩昕").Find(&u2, &u3)
	fmt.Printf("user:%#v\n", u2)

	//

	//3.更新
	//u.Name = "气密"
	//u.Age = 66
	//db.Debug().Save(&u)

	db.Debug().Model(&user).Update("name", "米琪")
}

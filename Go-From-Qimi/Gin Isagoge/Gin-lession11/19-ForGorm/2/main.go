package main

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

//截取自源代码
//type Model struct {
//	ID        uint `gorm:"primarykey"`
//	CreatedAt time.Time
//	UpdatedAt time.Time
//	DeletedAt DeletedAt `gorm:"index"`
//}

//定义模型
type UserSecond struct {
	gorm.Model   //内嵌gorm.model
	Name         string
	Age          sql.NullInt64 //零值类型
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
}

type Animal struct {
	ID       int    `gorm:"primary_key"`
	Name     string `gorm:"column:yourname"`
	Age      int
	BrithDay string
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/study-gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//db.AutoMigrate(&UserSecond{})
	//自动迁移
	//db.AutoMigrate(&UserInfo{})

	////添加
	//u1 := UserInfo{1, "气密", "男", "篮球"}
	//u2 := UserInfo{2, "mhx", "man", "football"}
	//db.Create(&u1)
	//db.Create(&u2)
	//
	//
	////查询
	//var u = new(UserInfo)
	//db.First(u)
	//fmt.Printf("%#v\n", u)
	//
	//var uu UserInfo
	//db.Find(&uu, "hobby=?", "football")
	//fmt.Printf("%#v\n", uu)
	//
	////更新
	//db.Model(&u).Update("hobby", "双色球")
	//fmt.Printf("%#v\n", u)
	//db.Model(&uu).Update("hobby", "asd")
	//fmt.Printf("%#v\n", uu)
	//
	////删除
	//db.Delete(&u)
	//db.Delete(&uu)

	db.AutoMigrate(&Animal{})

}
func (Animal) TableName() string {
	return "changename"
}

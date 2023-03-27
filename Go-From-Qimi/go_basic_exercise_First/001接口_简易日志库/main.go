package main

import (
	"fmt"
	"os"
	"time"
)

type Logger interface {
	consoleLog() //终端
	fileLog()    //文件
}

//用户结构体
type User struct {
	username string
	password string
}

//User实现方法
func (u User) consoleLog() {
	t := time.Now()
	fmt.Printf("用户创建成功！用户名：%s", u.username)
	fmt.Printf("用户创建成功！密码名：%s", u.password)
	fmt.Printf("创建完成时间：%d-%d-%d %d:%d:%d\n", t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
}

func (u User) fileLog() {
	t := time.Now()
	file, err := os.OpenFile("./"+u.username+".txt", os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := "用户创建成功！ 用户名为：" + fmt.Sprintf("%s\n", u.username) + "密码是：" + u.password + "\n" + fmt.Sprintf("创建完成时间：%d-%d-%d %d:%d:%d\n", t.Year(),
		t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	_, _ = file.WriteString(data)
	_ = file.Close()
}

// User中的字段初始化
func newUser(username, password string) User {
	return User{
		username: username,
		password: password,
	}
}

// 创建用户对象
func createUser() {
	var (
		username string
		password string
	)
	fmt.Print("请输入用户名：")
	_, err := fmt.Scan(&username)
	fmt.Print("请输入一个密码：")
	_, err = fmt.Scan(&password)
	if err != nil {
		fmt.Println("输入错误！！ERROR:", err)
	}
	u := newUser(username, password)
	u.consoleLog()
	u.fileLog()
}

func main() {
	createUser()
}

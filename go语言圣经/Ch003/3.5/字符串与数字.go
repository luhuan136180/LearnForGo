package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y)
	//z:=strconv.Itoa(x)
	fmt.Println(y, strconv.Itoa(x))

	fmt.Println(strconv.FormatInt(int64(x), 2))

	//	fmt.Sprintf返回一个格式化的字符串
	z := 1234
	s := fmt.Sprintf("Z=%b", z)
	fmt.Println(s)

	m := "123456"
	n, err := strconv.ParseInt(m, 10, 64)
	if err != nil {
		return
	}
	fmt.Println(n)
}

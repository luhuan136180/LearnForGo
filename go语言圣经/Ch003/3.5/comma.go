package main

import (
	"bytes"
	"fmt"
	"strings"
)

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

//练习 3.10： 编写一个非递归版本的comma函数，
//使用bytes.Buffer代替字符串链接操作。
func comma2(s string) string {
	byte := []byte(s)
	var buf bytes.Buffer

	if len(byte) <= 3 {
		return s
	}
	for i := 0; i < len(byte); i++ {
		if i != 0 && (len(byte)-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}

//练习 3.11： 完善comma函数，
//以支持浮点数处理和一个可选的正负号的处理。
func comma3(s string) string {
	index := strings.LastIndex(s, ".")
	if index == -1 {
		return comma2(s)
	} else {
		return comma2(s[:index]) + "." + comma2(s[index+1:])
	}
}
func main() {
	s := "1233.124"
	fmt.Println(comma3(s))
}

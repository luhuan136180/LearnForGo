package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "absxca"
	for val := range str {
		fmt.Println(val, " ")
	}
}

//将一个给定字符串 s 根据给定的行数 numRows ，
//以从上往下、从左到右进行 Z 字形排列。

//解析：以V字为一个循环，低谷只有一个元素，所以循环周期为n=2*numsRoes -2 ；
//用： x=i%n 确定i在循环中的位置
//行号y=min(x,n-x)  最主要是关注这个y的值，决定把它放在哪一行
func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	//仔细品这一步初始化切片的时候，不是make([]string,0)，
	//而是要几行就初始化几个空字符串，用来当行存储
	rows := make([]string, numRows)

	n := 2*numRows - 2

	for i, val := range s {
		x := i % n
		//
		rows[min(x, n-x)] += string(val)
	}
	return strings.Join(rows, "")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

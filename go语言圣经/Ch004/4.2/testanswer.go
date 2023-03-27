package main

import (
	"fmt"
	"unicode"
)

func main() {
	testArr := [5]int{0, 1, 2, 3, 4}
	reverseArray(&testArr)
	fmt.Println(testArr)
	a := rotate(testArr[:], 2)
	fmt.Println(a)

	b := []string{"tao", "tao", "taoshihan", "shi", "shi", "han"}
	emptyString(b)
	d := []byte("abc bcd wer  sdsd  taoshihan     de")
	e := emptyString2(d)
	fmt.Println(string(e))
	f := []byte("abc bcd wer  sdsd  taoshihan     de")
	reverse1(f)
	fmt.Println(string(f))
}

/*
练习 4.3： 重写reverse函数，使用数组指针代替slice。
*/
func reverseArray(s *[5]int) {
	i, j := 0, len(*s)-1
	for i < j {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
		i += 1
		j -= 1

	}
}

/*
练习 4.4： 编写一个rotate函数，通过一次循环完成旋转。
*/
func rotate(s []int, r int) []int {
	lens := len(s)
	//创建一个空的指定长度的slice
	res := make([]int, lens)
	for i := 0; i < lens; i++ {
		index := i + r
		if index >= lens {
			index = index - lens
		}
		res[i] = s[index]
	}
	return res
}

/*
练习 4.5：写一个函数在原地完成消除[]string中相邻重复的字符串的操作。
*/
func emptyString(s []string) []string {
	for i, j := 0, 1; j < len(s); i, j = i+1, j+1 {
		if s[i] == s[j] {
			copy(s[j:], s[j+1:])
			s = s[:len(s)-1]
			i, j = i-1, j-1 //重新比较i，j索引的字符串，回退
		}
	}
	return s
}

/*
练习 4.6： 编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考unicode.IsSpace）替换成一个空格返回
*/
func emptyString2(s []byte) []byte {
	index := 0
	num := len(s)
	for i := 0; i < num; i++ {
		index = i + 1
		num = len(s)
		if index >= num {
			break
		}
		if unicode.IsSpace(rune(s[i])) && unicode.IsSpace(rune(s[index])) {
			//结合remove函数
			copy(s[i:], s[index:])
			s = s[:len(s)-1]
			i--
		}
	}
	return s
}

/*
练习 4.7： 修改reverse函数用于原地反转UTF-8编码的[]byte。是否可以不用分配额外的内存？
*/
func reverse1(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

package main

import "fmt"

func main() {
	testArr := [5]int{0, 1, 2, 3, 4}
	reverse3(&testArr)
	fmt.Println(testArr)
	//a:=rotate(testArr[:],2)
	//fmt.Println(a)
	//
	//b:=[]string{"tao","taoshihan","shi","shi","han"}
	//emptyString(b)
	//d:=[]byte("abc bcd wer  sdsd  taoshihan     de")
	//e:=emptyString2(d)
	//fmt.Println(string(e))
	//f:=[]byte("abc bcd wer  sdsd  taoshihan     de")
	//reverse1(f)
	//fmt.Println(string(f))
}

func reverse3(s *[5]int) {
	i, j := 0, len(*s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
	}

}

/*
练习 4.5：写一个函数在原地完成消除[]string中相邻重复的字符串的操作。
*/
func emptyString3(s []string) []string {
	i := 0
	index := 0
	len := len(s)
	for _, val := range s {
		index = i + 1
		if index >= len {
			break
		}
		if val != s[index] {
			s[i] = val
			i++
		}

	}
	fmt.Println(s[:i])
	return s[:i]
}

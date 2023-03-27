package main

import (
	"fmt"
)

func main() {
	//s := "asda"
	//fmt.Println(&s)
	//t := 'a' //var t int32 = 'a'
	//fmt.Printf("t:%d", t)
	//fmt.Println()
	//s += ",hello"
	//q := s
	//fmt.Println(&s)
	//fmt.Println(&q)

	//s := "Hello, 世界"
	//fmt.Println(len(s))                    // "13"
	//fmt.Println(utf8.RuneCountInString(s)) // "9"  RuneCountInString类似于RuneCount，但它的输入是一个字符串。
	//buf := []byte("Hello, 世界")
	//fmt.Println("bytes =", len(buf))
	//fmt.Println("runes =", utf8.RuneCount(buf))
	//fmt.Println([]rune(s))

	//Go语言的range循环在处理字符串的时候，会自动隐式解码UTF8字符串。
	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	fmt.Println(string(1234567)) // "?"
}

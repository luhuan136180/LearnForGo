package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	l := "Hello, world.世界"
	for i, r := range l {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	fmt.Println(len(l))
	fmt.Println(utf8.RuneCountInString(l))

	//字符串和字节slice之间可以相互转换：
	s := "abc"
	b := []byte(s)
	s2 := string(b)

	fmt.Println(s)
	fmt.Println(s2)
	fmt.Println(b)
}

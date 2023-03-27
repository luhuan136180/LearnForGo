package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	//reader := strings.NewReader("GO语言中文网")
	//p := make([]byte, 6)
	//n, err := reader.ReadAt(p, 8)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%s,%d", p, n)
	//
	str := "GO语言中文网"
	s := "中文"
	index := strings.Index(str, s)
	if index < 0 {
		fmt.Println(index)
	}
	fmt.Println(utf8.RuneCountInString(str[:index]))
}

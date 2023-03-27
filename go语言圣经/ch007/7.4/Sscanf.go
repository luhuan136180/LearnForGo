package main

import "fmt"

func main() {
	s, t, u := "test123 "+
		"aaaaa", "", ""
	fmt.Sscan(s, &t)
	fmt.Println("s:", s)
	fmt.Println("t:", t) // t: test123 将s的内容传给t

	fmt.Sscanln(s, &t, &u)
	fmt.Println("s:", s)
	fmt.Println("t:", t) // t: test123 将s的内容传给t
	fmt.Println("u:", u) // t: test123 将s的内容传给t

	_, err := fmt.Sscanf(s, "test%s%s", &t, &u)
	fmt.Println("err:", err)
	fmt.Println("s:", s) // s: test123
	fmt.Println("t:", t) // t: 123 将t从s中去掉“test”后提取出来
	fmt.Println("u:", u)
}

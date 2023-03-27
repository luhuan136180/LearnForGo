package main

import "fmt"

func main() {
	p := new(int)
	*p = 1
	fmt.Println(p)
	fmt.Println(*p)
	*p = 2
	fmt.Println(p)
	//每次调用new函数都是返回一个新的变量的地址
	q := new(int)
	fmt.Println(q)
	a := 12
	fmt.Println(a)
}

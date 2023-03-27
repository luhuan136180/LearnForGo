package main

import "fmt"

type X string

//注：对x 必须进行 string（） 转换，否则会导致栈溢出
func (x X) String() string { return fmt.Sprintf("<%s>", string(x)) }
func main() {
	var x X
	x = "sadads"
	a := x.String()
	fmt.Println(a)

	//num := "123131"
	//s := fmt.Sprintf("<%s>", x)
	//fmt.Println(s)
}

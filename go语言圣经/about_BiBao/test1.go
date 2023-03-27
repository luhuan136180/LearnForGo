package main

import "fmt"

//闭包的延迟响应
func t1(x int) []func() {
	var fs []func()
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		fmt.Printf("outer val = %d, addr = %v\n", x+val, &val)
		fs = append(fs, func() {
			//这里传入的实际上是val的引用
			fmt.Printf("inner val = %d, addr = %v\n", x+val, &val)
		}) //这里只是声明，不会立刻执行，调用内部函数时才会真正执行
	}
	return fs
}
func main() {
	funcs := t1(11) //回到main中时，val已经 = 5了，此时fs中的各个隐匿函数还没有执行
	fmt.Println("declare done")
	for _, f := range funcs {
		f() //这里才会执行，val = 5，x = 11
	}
}

//
//out:
//outer val = 12, addr = 0xc0000ac058
//outer val = 13, addr = 0xc0000ac058
//outer val = 14, addr = 0xc0000ac058
//outer val = 16, addr = 0xc0000ac058
//declare done
//inner val = 16, addr = 0xc0000ac058
//inner val = 16, addr = 0xc0000ac058
//inner val = 16, addr = 0xc0000ac058
//inner val = 16, addr = 0xc0000ac058

package main

import "fmt"

//标准库提供了高度优化的bytes.Equal函数来判断两个字节型slice是否相等

//一个零值的slice等于nil。一个nil值的slice并没有底层数组。

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
func main() {
	//非nil值的slice的长度和容量也可以是0的
	var s []int    // len(s) == 0, s == nil
	s = nil        // len(s) == 0, s == nil
	s = []int(nil) // len(s) == 0, s == nil
	s = []int{}    // len(s) == 0, s != nil
	fmt.Println(s)
}

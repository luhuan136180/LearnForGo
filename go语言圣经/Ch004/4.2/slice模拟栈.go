package main

import "fmt"

//模拟栈
//stack = append(stack, v) // push v
//top := stack[len(stack)-html] // top of stack
//stack = stack[:len(stack)-html] // pop

//删除slice中间的某个元素并保存原有的元素顺序

func remove(slice []int, i int) []int {
	copy(slice[:i], slice[i+1:])
	return slice[:len(slice)-1]
}

func add(slice []int, index int, val int) []int {

	slice = append(slice[:index], append([]int{val}, slice[index:]...)...)
	return slice
}
func main() {
	s := []int{2, 3, 5, 6, 7, 8, 1}
	fmt.Println(remove(s, 2))
	fmt.Println("=====")
	fmt.Println(add(s, 4, 11))
}

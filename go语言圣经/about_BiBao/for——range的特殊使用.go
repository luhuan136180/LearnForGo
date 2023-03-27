package main

import "fmt"

func test() {
	arr := []int{1, 2, 3}
	newArr := []*int{}
	for _, v := range arr {
		newArr = append(newArr, &v)
		fmt.Printf("addr:%v\n", &v)
	}
	for _, v := range newArr {
		fmt.Println(*v)
	}
}

//out:
//addr:0xc0000a6010
//addr:0xc0000a6010
//addr:0xc0000a6010
//3
//3
//3
//正确写法
func main() {
	arr := []int{1, 2, 3}
	newArr := []*int{}
	for i, _ := range arr {
		newArr = append(newArr, &arr[i])
	}
	for _, v := range newArr {
		fmt.Println(*v)
	}
}

//out
//1
//2
//3

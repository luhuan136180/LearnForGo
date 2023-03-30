package main

import "fmt"

func main() {
	//s := []int{1, 2, 3, 4, 5, 6}

	s := []int{1, 2, 4, 7}

	s = Add(s, 0, 5)
	fmt.Println(s)
	s = Add(s, 1, 9)
	fmt.Println(s)
	s = Add(s, 6, 13)
	s = Delete(s, 2)
	fmt.Println(s)
	s = Delete(s, 0)
	fmt.Println(s)
	s = Delete(s, 4)
	fmt.Println(s)
}

func Add(s []int, index int, vaule int) []int {
	//do
	temp := []int{vaule}
	s = append(s[:index], append(temp, s[index:]...)...)
	return s
}

func Delete(s []int, index int) []int {
	//do

	s = append(s[:index], s[index+1:]...)
	return s
}

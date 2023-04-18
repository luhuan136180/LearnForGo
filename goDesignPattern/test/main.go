package main

import "fmt"

func main() {
	s := make([]int, 3, 4)

	fmt.Println(cap(s))
	news := myAppends(s)
	fmt.Println(news)
	s = append(s, 200)
	fmt.Println(s)
	fmt.Println(news)

}
func myAppends(s []int) []int {
	s = append(s, 100)
	return s
}

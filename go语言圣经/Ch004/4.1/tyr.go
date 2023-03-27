package main

import (
	"fmt"
)

func main() {
	s := [3]int{1, 2, 4}
	l := []int{1, 2, 3, 1}
	try(s, l)
	fmt.Println(s)
	fmt.Println(l)

}
func try(s [3]int, l []int) {
	s[2] = 82
	l[3] = 21
	fmt.Println(s)
	fmt.Println(l)

}

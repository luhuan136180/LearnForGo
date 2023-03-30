package main

import "fmt"

func main() {
	n := 4
	m := Solution(n)
	fmt.Println(m)
}

func Solution(n int) int64 {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return Solution(n-1) + Solution(n-2)
}

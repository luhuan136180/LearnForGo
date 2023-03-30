package main

import "fmt"

func main() {
	var n int
	for {

		fmt.Scan(&n)
		if n == 0 {
			break
		}
		arr := make([]int, n)
		fmt.Scanln(&arr)
		sum := 0
		for _, val := range arr {
			sum += val
		}
		fmt.Println(sum)
	}
}

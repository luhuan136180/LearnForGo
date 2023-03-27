package main

import "fmt"

func main() {
	ch := make(chan int, 32)
	go func() {
		ch <- searchByBing("golang")
	}()
	go func() {
		ch <- searchByGoogle("golang")
	}()
	go func() {
		ch <- searchByBaidu("golang")
	}()

	fmt.Println(<-ch)
}

func searchByBing(n string) int {
	return 1
}
func searchByGoogle(n string) int {
	return 1
}
func searchByBaidu(n string) int {
	return 1
}

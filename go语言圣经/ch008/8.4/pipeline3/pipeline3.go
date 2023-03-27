package main

import "fmt"

//chan<- int表示一个只发送int的channel
//<-chan int表示一个只接收int的channe
func Counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x

	}
	close(out)
}

//对一个只接收的channel调用close将是一个编译错误
func Squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}
func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go Counter(naturals)
	go Squarer(squares, naturals)
	printer(squares)
}

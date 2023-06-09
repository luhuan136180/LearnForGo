package main

import "fmt"

// 返回生成自然数序列的管道: 2, 3, 4, ...
func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()

	return ch
}

//管道过滤器: 删除能被素数整除的数
func PrimerFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for true {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func main() {
	ch := GenerateNatural() //
	for i := 0; i < 100; i++ {
		prime := <-ch
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimerFilter(ch, prime) // 基于新素数构造的过滤器
	}
}

package main

import (
	"sync"
	"sync/atomic"
)

var total2 uint64

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	var i uint64
	for i = 0; i <= 100; i++ {
		//atomic.AddUint64 函数调用保证了
		// total 的读取、更新和保存是一个原子操作，
		atomic.AddUint64(&total2, i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go worker(&wg)
	go worker(&wg)
	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
)

var total struct {
	sync.Mutex
	value int
}

//锁的原子操作-----初级

func Worker(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= 100; i++ {
		total.Lock()
		total.value += i
		total.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go Worker(&wg)
	go Worker(&wg)
	wg.Wait()

	fmt.Println(total.value)
}

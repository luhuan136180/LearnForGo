package main

import (
	"fmt"
	"sync"
	"time"
)

func worker2(wg *sync.WaitGroup, cancle chan bool) {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("hello")
		case <-cancle:
			return
		}
	}
}

func main() {
	cancle := make(chan bool)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker2(&wg, cancle)
	}

	time.Sleep(time.Second)
	close(cancle)
	wg.Wait()
}

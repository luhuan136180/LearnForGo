package main

func learn() {
	ch := make(chan int) // ch has type 'chan int'
	//
	var x int
	ch <- x  // a send statement
	x = <-ch // a receive expression in an assignment statement
	<-ch     // a receive statement; result is discarded

	close(ch)

	ch = make(chan int)    // unbuffered channel
	ch = make(chan int, 0) // unbuffered channel
	ch = make(chan int, 3) // buffered channel with capacity 3

}

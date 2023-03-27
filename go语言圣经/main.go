package main

func main() {
	//ch := make(chan int, 3)
	//x, m, n := 1, 3, 5
	//
	//go func() {
	//	ch <- x // a send statement
	//	ch <- m
	//	ch <- n
	//}()
	//close(ch)
	//y := <-ch // a receive expression in an assignment statement
	//fmt.Println(y)
	//for y = range ch {
	//	fmt.Println(y)
	//}

	//ch2 := make(chan []string)
	//l := []string{"123", "aaa", "bbb"}
	//var s []string
	//ch2 <- l
	//close(ch2)
	//s = <-ch2
	//fmt.Println(s)

	//for s = range ch2 {
	//	fmt.Println(s)
	//}

	ch := make(chan int)
	x := 3
	ch <- x
	y := <-ch
	panic(y)
}

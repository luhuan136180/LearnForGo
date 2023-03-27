package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

//func handleConn(c net.Conn) {
//	io.Copy(c, c) // NOTE: ignoring errors
//	c.Close()
//}

//func echo(c net.Conn, shout string, delay time.Duration) {
//	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
//	time.Sleep(delay)
//	fmt.Fprintln(c, "\t", shout)
//	time.Sleep(delay)
//	fmt.Fprintln(c, "\t", strings.ToLower(shout))
//}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)

	var wg sync.WaitGroup // number of working goroutines
	//abort := make(chan struct{})
	abort := make(chan string)
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-time.After(10 * time.Second):
				c.Close()
			case str := <-abort:
				wg.Add(1)
				go func(c net.Conn, shut string, delay time.Duration) {
					defer wg.Done()
					fmt.Fprintf(c, "\t", strings.ToUpper(shut))
					time.Sleep(delay)
					fmt.Fprintf(c, "\t", shut)
					time.Sleep(delay)
					fmt.Fprintf(c, "\t", strings.ToLower(shut))
				}(c, str, 1*time.Second)
			}
		}
	}()
	for input.Scan() {
		str := input.Text()
		if str == "exit" {
			break
		}
		if len(str) > 0 {
			abort <- str
		}

	}
	wg.Wait()
	c.Close()
}
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		coon, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(coon) //针对复数个tcp链接开启携程
	}
}

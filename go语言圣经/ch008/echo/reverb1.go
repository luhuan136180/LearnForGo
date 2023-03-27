package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

//func handleConn(c net.Conn) {
//	io.Copy(c, c) // NOTE: ignoring errors
//	c.Close()
//}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c) //从连接中读取信息，及从client端读取信息
	for input.Scan() {
		//针对复数个输入开启携程
		go echo(c, input.Text(), 1*time.Second)
	}
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

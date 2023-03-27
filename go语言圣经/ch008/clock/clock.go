package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		coon, err := listener.Accept() //平时阻塞，在接收到链接请求时创建链接，返回一个conn
		//net.Conn对象来表示这个连接。
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(coon)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

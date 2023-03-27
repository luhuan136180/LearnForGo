package main

import (
	"io"
	"log"
	"net"
	"os"
)

//一个TCP连接有读和写两个部分，
//可以使用CloseRead和CloseWrite方法分别关闭它们。
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdin, conn)
		log.Println("done")
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	//类型断言，调用*net.TCPConn的方法CloseWrite()只关闭TCP的写连接
	cw := conn.(*net.TCPConn)
	cw.CloseWrite()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	//换成常规的net链接方法了
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("net.Dial", err)
	}
	//NewClientWithCodec 类似  NewClient  ；
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}

package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	//Dial在指定的网络和地址与RPC服务端连接。
	//返回一个rpcclient 和一个 err
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string

	//client.Call 调用具体的 RPC 方法
	//第一个参数是用点号连接的 RPC 服务名字和方法名字，
	//第二和第三个参数分别我们定义 RPC 方法的两个参数
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}

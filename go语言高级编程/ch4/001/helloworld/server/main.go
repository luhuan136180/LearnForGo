package main

import (
	"log"
	"net"
	"net/rpc"
)

//将HelloService类型的对象注册为一个RPC服务
type HelloService struct {
}

// Go 语言的 RPC 规则：rpc服务的方法只能有两个可序列化的参数，
//其中第二个参数是指针类型，
//并且返回一个 error 类型，同时必须是公开的方法。
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello " + request
	return nil
}

func main() {
	//调用会将对象类型中所有满足 RPC 规则的对象方法注册为 RPC 函数，
	//所有注册的方法会放在 “HelloService” 服务空间之下
	rpc.RegisterName("HelloService", new(HelloService))

	//建立一个唯一的 TCP 连接
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)

	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}
	//ServeConn() 在单个连接上执行server
	rpc.ServeConn(conn)
}

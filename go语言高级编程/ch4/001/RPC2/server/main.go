package main

import (
	"log"
	"net"
	"net/rpc"
)

//服务的名字
const HelloServiceName = "path/to/pkg.HelloService"

//服务实现的详细方法列表
type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}

//注册该类型服务的函数
func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

//一个实现
type HelloService struct {
}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello " + request
	return nil
}

func main() {

	//rpc.RegisterName("HelloService", new(HelloService))

	RegisterHelloService(new(HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)

	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeConn(conn)
	}
}

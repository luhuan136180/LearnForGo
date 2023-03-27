package main

import (
	"net"
	"net/rpc"
	"time"
)

type HelloService struct {
}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello " + request
	return nil
}

//启动反向 RPC 服务的代码:
//反向 RPC 的内网服务将不再主动提供 TCP 监听服务，
//而是首先主动连接到对方的 TCP 服务器。
//然后基于每个建立的 TCP 连接向对方提供 RPC 服务。
func main() {
	//rcvr不是一个导出类型的值，或者该类型没有满足要求的方法，Register会返回错误
	rpc.Register(new(HelloService))
	//客户端可以使用格式为"Type.Method"的字符串访问这些方法，其中Type是rcvr的具体类型。
	for {
		//Listen函数创建的服务端：
		//
		//Dial函数和服务端建立连接：
		conn, _ := net.Dial("tcp", "localhost:1234")
		//Conn接口代表通用的面向流的网络连接。多个线程可能会同时调用同一个Conn的方法。
		if conn == nil {
			time.Sleep(time.Second)
			continue
		}
		rpc.ServeConn(conn)
		conn.Close()
	}
}

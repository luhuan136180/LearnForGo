package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {
	conn    net.Conn
	isLogin bool
}

//基于上下文信息，我们可以方便地为 RPC 服务增加简单的登陆状态的验证：
func (p *HelloService) Login(request string, reply *string) error {
	if request != "user:password" {
		return fmt.Errorf("auth failed")
	}
	log.Println("login ok")
	p.isLogin = true
	return nil
}
func (p *HelloService) Hello(request string, reply *string) error {
	if !p.isLogin {
		return fmt.Errorf("please login")
	}
	*reply = "hello " + request + ",from" + p.conn.RemoteAddr().String()
	return nil
}

func main() {
	//返回在一个本地网络地址laddr上监听的Listener。
	//网络类型参数net必须是面向流的网络：
	listener, err := net.Listen("tcp", ":1236")
	//Listener是一个用于面向流的网络协议的公用的网络监听器接口。
	//多个线程可能会同时调用一个Listener的方法。
	if err != nil {
		log.Fatal("ListenTCP error :", err)
	}

	for {
		//// Accept等待并返回下一个连接到该接口的连接
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error :", err)
		}

		go func() {
			defer conn.Close()

			p := rpc.NewServer()
			p.Register(&HelloService{
				conn: conn,
			})
			p.ServeConn(conn)
		}()
	}
}

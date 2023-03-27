package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {
}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello " + request
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP . error", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		//ServeCodec类似ServeConn，但使用指定的编解码器，
		//以编码请求主体和解码回复主体。
		//传递一个ServerCodec接口
		//具体值是一个：针对服务端的 json 编解码器。
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}

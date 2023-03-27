package main

import (
	hello "Gobook3/ch4/004/gRPC-Basic"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type HelloServiceImpl struct {
}

func (p *HelloServiceImpl) Hello(ctx context.Context,
	args *hello.String) (*hello.String, error) {
	reply := &hello.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}
func main() {
	grpcService := grpc.NewServer()
	hello.RegisterHelloServiceServer(grpcService, new(HelloServiceImpl))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcService.Serve(lis)
}

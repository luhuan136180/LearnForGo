package main

import (
	"Gobook3/ch4/003/RPC-Watch/KV"
	"log"
	"net"
	"net/rpc"
)

func main() {
	rpc.RegisterName("KVStoreService", KV.NewKVStoreService())
	listner, err := net.Listen("tcp", ":1235")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	conn, err := listner.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}
	rpc.ServeConn(conn)
}

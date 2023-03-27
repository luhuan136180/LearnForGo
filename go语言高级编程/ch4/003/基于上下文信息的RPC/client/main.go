package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1236")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply string
	err = client.Call("HelloService.Login", "user:password", &reply)
	if err != nil {
		log.Fatal("login failed:", err)
	} else {
		fmt.Println("login success...")
	}
	err = client.Call("HelloService.Hello", "hello-aa", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}

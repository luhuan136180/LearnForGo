package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

func doClientWork(client *rpc.Client) {
	go func() {
		var keyChanged string

		err := client.Call("KVStoreService.Watch", 30, &keyChanged)
		fmt.Println("===")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("watch:", keyChanged)
	}()
	time.Sleep(time.Second * 2)
	err := client.Call("KVStoreService.Set", [2]string{"abc", "abc-value3"},
		new(struct{}))
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 3)
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1235")
	if err != nil {
		log.Fatal("dailing :", err)
	}
	doClientWork(client)

}

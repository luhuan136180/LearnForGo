package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex //互斥锁
var count int

func main() {
	http.HandleFunc("/", handler2)
	http.HandleFunc("/count", counter)

	//log.Fatal  用于监视，出现err时直接退出
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
func handler2(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q \n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

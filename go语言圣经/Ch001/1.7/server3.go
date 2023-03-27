package main

import (
	"fmt"
	"log"
	"net/http"
)

//var mu sync.Mutex
//var count int

func main() {
	http.HandleFunc("/", handler3)
	//http.HandleFunc("/count", counter3)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

// counter echoes the number of calls so far.
//func counter3(w http.ResponseWriter, r *http.Request) {
//	mu.Lock()
//	//Fprintf根据format参数生成格式化的字符串并写入w。
//	fmt.Fprintf(w, "Count %d\n", count)
//	mu.Unlock()
//}

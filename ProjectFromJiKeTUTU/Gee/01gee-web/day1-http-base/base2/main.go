package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct {
}

//正确的路由映射方法选择器
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		//fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
		geturl(w, req)
	case "/hello":
		//for k, v := range req.Header {
		//	fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		//}
		getheader(w, req)
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}
func geturl(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}
func getheader(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}

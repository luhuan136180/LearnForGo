package main

import (
	"fmt"
	"net/http"
)

func home() {

}
func SignUp(w http.ResponseWriter, r *http.Request) {
	req := &signUpReq{}
	ctx := &Context{
		W: w,
		R: r,
	}
	err := ctx.ReadJson(req)
	if err != nil {
		fmt.Fprintf(w, "err:%v", err)
		return
	}

}
func main() {
	server := NewHttpServer("test-server")
	server.Route("/", home)
	server.Route("/user", user)
	server.Start(":8080")
}

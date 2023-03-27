package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func echo(w http.ResponseWriter, r *http.Request) {
	msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("echo error"))
		return
	}

	writelen, err := w.Write(msg)
	if err != nil || writelen != len(msg) {
		log.Println(err, "write len:", writelen)
	}
}
func main() {
	http.HandleFunc("/", echo)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

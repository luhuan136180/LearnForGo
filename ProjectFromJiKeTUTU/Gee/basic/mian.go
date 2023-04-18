package basic

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	//http.HandlerFunc("/count",counter)
	http.ListenAndServe("localhost:8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q \n", r.URL.Path)
}

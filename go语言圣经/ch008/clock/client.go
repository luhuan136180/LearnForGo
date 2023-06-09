package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	coon, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer coon.Close()

	mustCopy(os.Stdout, coon)
}
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

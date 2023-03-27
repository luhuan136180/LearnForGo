package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	io.Copy(os.Stdout, os.Stdin)
	fmt.Println("Got EOF -- bye")
	io.CopyN(os.Stdout, strings.NewReader("Go语言中文网"), 8)
}

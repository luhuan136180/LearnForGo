package main

import (
	"bufio"
	"fmt"
	"strings"
	"time"
)

//func main() {
//	reader := bufio.NewReaderSize(strings.NewReader("http://studygolang.com.\t It is the home of gophers"), 14)
//	go Peek(reader)
//
//	go reader.ReadBytes('\t')
//	time.Sleep(1e8)
//}
//
//func Peek(reader *bufio.Reader) {
//	line, _ := reader.Peek(14)
//	fmt.Printf("%s\n", line)
//	fmt.Printf("%s\n", line)
//}

func main() {
	reader := bufio.NewReaderSize(strings.NewReader("http://studygolang.com.\t It is the home of gophers"), 14)
	go Peek(reader)
	//time.Sleep(1e8)
	go reader.ReadBytes('\t')

	time.Sleep(1e8)
}

func Peek(reader *bufio.Reader) {
	line, _ := reader.Peek(14)
	fmt.Printf("%s\n", line)
	// time.Sleep(1)
	fmt.Printf("%s\n", line)
}

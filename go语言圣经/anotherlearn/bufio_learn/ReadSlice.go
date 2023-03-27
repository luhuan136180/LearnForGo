package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))

	line, _ := reader.ReadSlice('\n')
	fmt.Printf("the line :%s\n", line)

	n, _ := reader.ReadSlice('\n')
	fmt.Printf("the line:%s\n", line)
	//fmt.Printf("the line:%s\n", n)

	fmt.Println(string(n))

	reader2 := bufio.NewReaderSize(strings.NewReader("http://studygolang.com"), 16)
	line, err := reader2.ReadSlice('\n')
	fmt.Printf("line:%s\terror:%s\n", line, err)
	line, err = reader2.ReadSlice('\n')
	fmt.Printf("line:%s\terror:%s\n", line, err)

}

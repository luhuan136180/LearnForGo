package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//scanner := bufio.NewScanner(os.Stdin)
	//for scanner.Scan() {
	//	fmt.Println(scanner.Text()) // Println will add back the final '\n'
	//}
	//if err := scanner.Err(); err != nil {
	//	fmt.Fprintln(os.Stderr, "reading standard input:", err)
	//}
	//ScanWords
	const input = "This is The Golang Standard library.\n Welcome you!"
	scanner := bufio.NewScanner(strings.NewReader(input))
	//通过调用 scanner.Split(bufio.ScanWords)
	//来更改 split 函数
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)

	}
	fmt.Println(count)

	//Scanner
	file, err := os.Create("scanner.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString("http://studygolang.com.\nIt is the home of gophers.\nIf you are studying golang, welcome you!")
	// 将文件 offset 设置到文件开头
	//file.Seek(0, os.)

	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

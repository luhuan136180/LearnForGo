package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//创建时初始化，并赋值，数值类=0，字符类为空
	counts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		//counts[input.Text()]++
		line := input.Text()
		counts[line] = counts[line] + 1
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

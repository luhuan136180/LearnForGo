package main

import (
	"bufio"
	"fmt"
	"os"
)

type LnFile struct {
	Count    int
	FileName string
}

func main() {
	counts := make(map[string]*LnFile)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n.Count > 1 {

			fmt.Printf("%d\t%v\t%s\n", n.Count, n.FileName, line)
		}
	}
}

func countLines(f *os.File, counts map[string]*LnFile) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if _, ok := counts[input.Text()]; ok {
			//已经录入
			counts[input.Text()].Count++
		} else {
			//需要初始化
			counts[input.Text()] = new(LnFile)
			counts[input.Text()].Count++
			counts[input.Text()].FileName = f.Name()
		}

	}
}

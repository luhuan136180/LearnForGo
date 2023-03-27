package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		//func HasPrefix(s, prefix string) bool
		//判断s是否有前缀字符串prefix。
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}
		//b, err := ioutil.ReadAll(resp.Body)

		//func Copy(dst Writer, src Reader) (written int64, err error)
		//将src的数据拷贝到dst
		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:reading %s:%v\n", url, err)
			os.Exit(1)
		}
		fmt.Println("________-------_________")
		fmt.Printf("%d\n", b)
		fmt.Printf("@@@%s", resp.Status)
	}
}

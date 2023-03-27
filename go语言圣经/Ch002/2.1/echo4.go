package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("S", " ", "separator")

func main() {
	//从os.Args[html:]中解析注册的flag
	//必须在所有flag都注册好而未访问其值时执行。
	flag.Parse()
	//func Join(a []string, sep string) string
	//将一系列字符串连接为一个字符串，之间用sep来分隔
	fmt.Println(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Println("$")
	}
}

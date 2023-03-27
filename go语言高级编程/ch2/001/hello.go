package main

import "C"

//通过 import "C" 语句启用 CGO 特性

//func main() {
//	println("hello cgo")
//}

func main() {
	C.puts(C.CString("Hello,world\n"))
}

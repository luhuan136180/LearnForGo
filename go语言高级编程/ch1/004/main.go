package main

import (
	"fmt"
	"os"
)

//var (
//	a io.ReadCloser = (*os.File)(f) // 隐式转换, *os.File 满足 io.ReadCloser 接口
//	b io.Reader     = a             // 隐式转换, io.ReadCloser 满足 io.Reader 接口
//	c io.Closer     = a             // 隐式转换, io.ReadCloser 满足 io.Closer 接口
//	d io.Reader     = c.(io.Reader) // 显式转换, io.Closer 不满足 io.Reader 接口
//)

func main() {
	//闭包对捕获的外部变量并不是传值方式访问，而是以引用的方式访问
	for i := 0; i < 3; i++ {
		defer func() { println(i) }()
	}
	//，在 for 迭代语句中，每个 defer
	//语句延迟执行的函数引用的都是同一个 i 迭代变量，
	//在循环结束后这个变量的值为 3，因此最终输出的都是3

	//在循环体内部再定义一个局部变量，这样每次迭代 defer
	//语句的闭包函数捕获的都是不同的变量，
	//这些变量的值对应迭代时的值。
	for i := 0; i < 3; i++ {
		i := i //定义一个循环体内的局部变量
		defer func() { println(i) }()
	}

	//方法转换为原始函数

	//// 不依赖具体的文件对象
	//// func CloseFile(f *File) error
	//var CloseFile = (*File).Close
	//
	//// 不依赖具体的文件对象
	//// func ReadFile(f *File, offset int64, data []byte) int
	//var ReadFile = (*File).Read
	//
	//// 文件处理
	//f, _ := OpenFile("foo.dat")
	//ReadFile(f, 0, data)
	//CloseFile(f)

	fmt.Fprintln(os.Stdout, Upperstring("hello,world"))
}

type Upperstring string

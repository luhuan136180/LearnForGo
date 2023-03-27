package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

func main() {
	var a = [...]int{1, 2, 3}
	var b = &a
	fmt.Println(a[0], a[1])
	fmt.Println(b[1], b[0])

	for i, v := range b {
		fmt.Println(i, v)
	}
	fmt.Println("=================")
	for index, val := range a {
		fmt.Println(index, ":", val)
	}

	//
	// 字符串数组
	//var s1 = [2]string{"hello", "world"}
	//var s2 = [...]string{"你好", "世界"}
	//var s3 = [...]string{1: "世界", 0: "你好"}
	//
	//// 结构体数组
	//var line1 [2]image.Point
	//var line2 = [...]image.Point{image.Point{X: 0, Y: 0}, image.Point{X: 1, Y: 1}}
	//var line3 = [...]image.Point{{0, 0}, {1, 1}}
	//
	//// 图像解码器数组
	//var decoder1 [2]func(io.Reader) (image.Image, error)
	//var decoder2 = [...]func(io.Reader) (image.Image, error){
	//	png.Decode,
	//	jpeg.Decode,
	//}
	//
	//// 接口数组
	//var unknown1 [2]interface{}
	//var unknown2 = [...]interface{}{123, "你好"}
	//
	//// 管道数组
	//var chanList = [2]chan int{}
	//
	////	空数组
	//var d [0]int       // 定义一个长度为 0 的数组
	//var e = [0]int{}   // 定义一个长度为 0 的数组
	//var f = [...]int{} // 定义一个长度为 0 的数组

	//	空数组的应用场景
	c1 := make(chan [0]int)
	go func() {
		fmt.Println("c1")
		c1 <- [0]int{}
	}()
	<-c1
	close(c1)
	//优化
	c2 := make(chan struct{})
	go func() {
		fmt.Println("c2")
		c2 <- struct{}{}
	}()
	<-c2
	close(c2)

	s := "hello, world"
	hello := s[:5]
	world := s[7:]

	s1 := "hello, world"[:5]
	s2 := "hello, world"[7:]

	fmt.Println(hello)
	fmt.Println(world)
	fmt.Println(s1)
	fmt.Println(s2)

}

//for range 对字符串的迭代模拟实现
func forOnString(s string, forBody func(i int, r rune)) {
	for i := 0; len(s) > 0; {
		r, size := utf8.DecodeRuneInString(s)
		forBody(i, r)
		s = s[size:]
		i += size
	}
}

//[]byte(s) 转换模拟实现
func str2bytes(s string) []byte {
	p := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		p[i] = c
	}
	return p
}

//string(bytes) 转换模拟实现
func bytesStr(s []byte) (p string) {
	data := make([]byte, len(s))
	for i, c := range s {
		data[i] = c
	}

	hdr := (*reflect.StringHeader)(unsafe.Pointer(&p))
	hdr.Data = uintptr(unsafe.Pointer(&data[0]))
	hdr.Len = len(s)

	return p
}

//[]rune(s) 转换模拟实现
func str2Runes(s string) []rune {
	var p []int32
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		p = append(p, int32(r))
		s = s[size:]
	}
	return []rune(p)
}

//string(runes) 转换模拟实现
func runes2string(s []int32) string {
	var p []byte
	buf := make([]byte, 3)
	for _, r := range s {
		n := utf8.EncodeRune(buf, r)
		p = append(p, buf[:n]...)
	}
	return string(p)
}

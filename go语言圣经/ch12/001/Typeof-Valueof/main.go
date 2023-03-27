package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	t := reflect.TypeOf(3)
	fmt.Println(t.String())
	fmt.Println(t)
	//Typeof函数：总是返回具体的类型
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w)) // "*os.File"
	//reflect.Type 接口是满足 fmt.Stringer 接口的
	fmt.Printf("%T\n", 3)

	//
	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Printf("%v\n", v)
	fmt.Println(v.String()) // NOTE: "<int Value>"
	//
	m := v.Type()
	//对 Value 调用 Type 方法将返回具体类型所对应的 reflect.Type：
	fmt.Printf(m.String())

	//reflect.ValueOf 的逆操作是 reflect.Value.Interface 方法
	v = reflect.ValueOf(3) // a reflect.Value
	x := v.Interface()     // an interface{}
	i := x.(int)           // an int
	fmt.Printf("%d\n", i)  // "3"

	//

}

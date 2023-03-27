package main

import (
	"fmt"
	"reflect"
)

func Add(a, b int) int {
	return a + b
}

func main() {
	v := reflect.ValueOf(Add)
	fmt.Println(v) //0x3dfc40
	if v.Kind() != reflect.Func {
		return
	}

	t := v.Type()
	fmt.Println(t) //func(int, int) int
	//reflect.rtype.NumIn 获取函数的入参个数；
	argv := make([]reflect.Value, t.NumIn())
	fmt.Println(argv) //[<invalid Value> <invalid Value>]
	for i := range argv {
		//fmt.Printf("%d-type:%v-kind:%d\n", i, t.In(i), t.In(i).Kind())  :0-type:int-kind:2
		//fmt.Println("reflect.ValueOf(i):", reflect.ValueOf(i))
		//// In()返回函数类型的第i个输入参数的类型。
		//
		////如果类型的Kind不是Func，它会恐慌。
		//
		////如果i不在[0,NumIn())的范围内，则会出现panic。
		//
		//// Kind() returns the specific kind of this type.
		if t.In(i).Kind() != reflect.Int {
			return
		}
		argv[i] = reflect.ValueOf(i)
	}
	//fmt.Println("===")
	//fmt.Println(argv)
	result := v.Call(argv)
	if len(result) != 1 || result[0].Kind() != reflect.Int {
		return
	}
	fmt.Println(result[0].Int()) // #=> 1
}

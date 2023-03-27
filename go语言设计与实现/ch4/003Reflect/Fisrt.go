package main

import (
	"fmt"
	"reflect"
)

//空接口
type a interface {
}

//实现了error接口
type CustomError struct {
}

func (*CustomError) Error() string {
	return ""
}
func main() {
	//suthor := "mhx"
	//t := reflect.TypeOf(suthor)
	//v := reflect.ValueOf(&suthor)
	////type的String函数自动调用
	//fmt.Println("typeof author", t)
	//fmt.Println("valueof author", v)
	//
	//fmt.Println("suthor=", suthor)
	////调用 reflect.Value.Elem 获取指针指向的变量；
	////调用 reflect.Value.SetInt 更新变量的值：
	//v.Elem().SetString("mhx====")
	//fmt.Println("suthor:", suthor)

	//如何判断一个类型是否实现了某个接口
	//
	//获得接口类型需要通过以下方式：reflect.TypeOf((*<interface>)(nil)).Elem()
	typeOfError := reflect.TypeOf((*error)(nil)).Elem()
	customErrorPtr := reflect.TypeOf(&CustomError{})
	customError := reflect.TypeOf(CustomError{})

	fmt.Println(typeOfError)
	fmt.Println(customErrorPtr)
	fmt.Println(customError)

	fmt.Println(customErrorPtr.Implements(typeOfError)) // #=> true
	fmt.Println(customError.Implements(typeOfError))    // #=> false

}

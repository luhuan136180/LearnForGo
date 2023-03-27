package main

import "fmt"

//适配的目标
type V5 interface { //一个接口
	Use5V()
}

//业务类，依赖V5接口
type Phone struct {
	v V5
}

//业务类，依赖V5接口
func NewPhone(v V5) *Phone {
	return &Phone{v}
}

func (p *Phone) Charge() {
	fmt.Println("Phone进行充电...")
	p.v.Use5V()
}

//被适配的角色，适配者
type V220 struct{} //需要5v，但只有220v——所以需要适配器

func (v *V220) Use220V() {
	fmt.Println("使用220V的电压")
}

//适配器实例类
type Adapter struct {
	v220 *V220
}

func (a *Adapter) Use5V() {
	fmt.Println("使用适配器进行充电")

	//调用适配者的方法
	a.v220.Use220V()
}

func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{v220}
}

func main() {
	iphone := NewPhone(NewAdapter(new(V220)))

	iphone.Charge()
}

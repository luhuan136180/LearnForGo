package main

import "fmt"

//   装饰器模式偏重对原对象功能的扩展，扩展后的对象仍是是对象本身；
//然而代理模式偏重因自己无法完成或无需关心，需要他人干涉事件流程，
//更多的是对对象的控制（代理使客户端不需要知道实现类是什么，
//怎么做的，而客户端只需知道代理即可，即将客户端与实现类解耦）

// ---------- 抽象层 ----------
//抽象的构件
type Phone interface {
	Show()
}

//抽象装饰器基础类（该类本应该为interface，
//但是Golang interface语法不可以有成员属性
type Decorator struct {
	phone Phone
}

//装饰器本身也要实现它要装饰的接口
func (d *Decorator) Show() {
	//空
}

//实现层----------------
// 具体的构件
type HuaWei struct{}

func (hw *HuaWei) Show() {
	fmt.Println("秀出了HuaWei手机")
}

type XiaoMi struct{}

func (xm *XiaoMi) Show() {
	fmt.Println("秀出了XiaoMi手机")
}

//具体的装饰器
type MoDecorator struct {
	Decorator //继承基础装饰器类(主要继承Phone成员属性)
}

func (md *MoDecorator) Show() {
	md.phone.Show()      //调用被装饰构件的原方法
	fmt.Println("贴膜的手机") //装饰额外的方法
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{Decorator{phone}}
}

type KeDecorator struct {
	Decorator //继承基础装饰器类(主要继承Phone成员属性)
}

func (kd *KeDecorator) Show() {
	kd.phone.Show()
	fmt.Println("手机壳的手机") //装饰额外的方法
}

func NewKeDecorator(phone Phone) Phone {
	return &KeDecorator{Decorator{phone}}
}

func main() {
	var huawei Phone //裸机华为手机
	huawei = new(HuaWei)
	huawei.Show() //调用原构件方法

	fmt.Println("---------")

	//用贴膜装饰器装饰，得到新功能构件
	var moHuawei Phone
	moHuawei = NewMoDecorator(huawei) //通过HueWei ---> MoHuaWei
	moHuawei.Show()                   //调用装饰后新构件的方法

	fmt.Println("---------")
	var keHuawei Phone
	keHuawei = NewKeDecorator(huawei) //通过HueWei ---> KeHuaWei
	keHuawei.Show()

	fmt.Println("---------")
	var keMoHuaWei Phone
	keMoHuaWei = NewMoDecorator(keHuawei) //通过KeHuaWei ---> KeMoHuaWei
	keMoHuaWei.Show()
}

package main

import "fmt"

//实现  开闭原则
//抽象的银行业务员
type AbstractBanker interface {
	DoBusi() //抽象的处理业务接口
}

//存款的业务员
type SaveBanker struct {
	//AbstractBanker
}

func (sb *SaveBanker) DoBusi() {
	fmt.Println("进行了存款")
}

//转账的业务员
type TransferBanker struct {
	//AbstractBanker
}

func (tb *TransferBanker) DoBusi() {
	fmt.Println("进行了转账")
}

//支付的业务员
type PayBanker struct {
	//AbstractBanker
}

func (pb *PayBanker) DoBusi() {
	fmt.Println("进行了支付")
}

type FundBanker struct {
}

func (fb *FundBanker) DoBusi() {

}

//func main() {
//	sb := &SaveBanker{}
//	sb.DoBusi()
//
//	//进行转账
//	tb := &TransferBanker{}
//	tb.DoBusi()
//
//	//进行支付
//	pb := &PayBanker{}
//	pb.DoBusi()
//}

//实现架构层(基于抽象层进行业务封装-针对interface接口进行封装)
func BankerBusiness(banker AbstractBanker) {
	//通过接口来向下调用，(多态现象)
	banker.DoBusi()
}
func main() {
	//原理：将具体是实现类用接口封装，让函数自己判断并调用相应的具体实现类型

	//进行存款
	BankerBusiness(&SaveBanker{})
	//进行存款
	BankerBusiness(&TransferBanker{})
	//进行存款
	BankerBusiness(&PayBanker{})
}

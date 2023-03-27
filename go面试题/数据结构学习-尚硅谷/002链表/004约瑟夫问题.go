package main

import "fmt"

type Boy struct {
	No   int
	Next *Boy
}

func AddBoy(num int) *Boy {
	first := &Boy{} //空节点
	helper := &Boy{}
	if num < 1 {
		fmt.Println("num有误")
		return first
	}
	for i := 1; i <= num; i++ {
		//
		boy := &Boy{
			No: i,
		}
		if i == 1 {
			first = boy
			helper = boy
			helper.Next = first
		} else {
			helper.Next = boy
			helper = boy
			helper.Next = first //
		}

	}
	return first
}

func SHowBoy(first *Boy) {
	if first.Next == nil {
		fmt.Println("没有·蒜素")
		return
	}
	curBoy := first
	for {
		//
		fmt.Printf("小孩编号=%d ->", curBoy.No)
		if curBoy.Next == first {
			break
		}
		curBoy = curBoy.Next
	}

}

func PlayGame(first *Boy, StartNum int, countNo int) {
	if first.Next == nil {
		fmt.Println("空的链表，没有小孩")
		return
	}
	tail := first
	//将辅助接点转移到first的前节点上
	for {
		if tail.Next == first {
			break
		}
		tail = tail.Next
	}

	//将first节点移动到起始节点
	for i := 1; i <= StartNum; i++ {
		first = first.Next
		tail = tail.Next
	}

	fmt.Println()

	//开始报数
	for {
		for i := 1; i < countNo-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("小孩编号为%d 出圈 \n", first.No)
		first = first.Next
		tail.Next = first
		if tail == first {
			break
		}
	}
	fmt.Printf("小孩小孩编号为%d 出圈 \n", first.No)
}

func main() {

	first := AddBoy(500)
	//显示
	SHowBoy(first)
	PlayGame(first, 20, 31)

}

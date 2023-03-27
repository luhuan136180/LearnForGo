package main

import "fmt"

type CatNode struct {
	no   int //猫猫的编号
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	//
	if head.next == nil { //链表为空
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head
		fmt.Println(newCatNode, "加入到环形的链表")
		return
	}

	//
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	temp.next = newCatNode
	newCatNode.next = head
}

func ListCricleLink(head *CatNode) {
	if head.next == nil {
		fmt.Println("空空如也的环形链表...")
		return
	}
	temp := head
	for {
		fmt.Printf("猫的信息：[id=%d,name = %s]->", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

func DelCatNode(head *CatNode, id int) *CatNode {
	temp := head
	helper := head
	if temp.next == nil {
		fmt.Println("这是一个空的环形链表，不能删除")
		return head
	}
	if temp.next == head {
		if temp.no == id {
			temp.next = nil
		}
		return head
	}
	//将 helper 定位到链表最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}
	//如果有两个包含两个以上结点
	flag := true
	for {
		if temp.next == head { //如果到这来，说明我比较到最后一个【最后一个还没比较】
			break
		}
		if temp.no == id {
			if temp == head { //说明删除的是头结点
				head = head.next
			}
			//恭喜找到., 我们也可以在直接删除
			helper.next = temp.next
			fmt.Printf("猫猫=%d\n", id)
			flag = false
			break
		}
		temp = temp.next     //移动 【比较】
		helper = helper.next //移动 【一旦找到要删除的结点 helper】
	}
	//这里还有比较一次
	if flag { //如果 flag 为真，则我们上面没有删除
		if temp.no == id {
			helper.next = temp.next
			fmt.Printf("猫猫=%d\n", id)
		} else {
			fmt.Printf("对不起，没有 no=%d\n", id)
		}
	}
	return head
}

package main

import "fmt"

type HeroNOde struct {
	no       int
	name     string
	nickname string
	next     *HeroNOde
}

func InsertHeroNode(head *HeroNOde, newHeroNode *HeroNOde) {
	//思路
	//1. 先找到该链表的最后这个结点
	//2. 创建一个辅助结点[跑龙套, 帮忙]
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newHeroNode
}

func InsertHeroNodeByNo(head *HeroNOde, newHeroNode *HeroNOde) {
	temp := head
	flag := true
	for {
		if temp.next == nil {
			break
		} else if temp.next.no >= newHeroNode.no {
			break
		} else if temp.next.no == newHeroNode.no {
			flag = false
			break
		}
		temp = temp.next

	}
	if !flag {
		fmt.Println("对不起，已经存在no=", newHeroNode)
		return
	} else {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
	}

}

//显示链表的所有结点信息
func ListHeroNode(head *HeroNOde) {
	temp := head
	if temp.next == nil {
		fmt.Println("没有")
		return
	}
	for {
		fmt.Printf("[%d,%s,%s]=>", temp.next.no, temp.next.name, temp.next.nickname)
		temp = temp.next

		if temp.next == nil {
			break
		}
	}
}

func DelHeroNode(head *HeroNOde, id int) {
	temp := head
	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			//找到了
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next
	} else {
		fmt.Printf("没有找到要删除的id为：%v 的节点", id)
	}
}
func main() {
	head := &HeroNOde{}

	hero := &HeroNOde{
		no:       1,
		name:     "songjiang",
		nickname: "jishiyu",
	}
	hero2 := &HeroNOde{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}
	InsertHeroNode(head, hero)
	InsertHeroNode(head, hero2)

	ListHeroNode(head)
	fmt.Println()
	DelHeroNode(head, 1)
	ListHeroNode(head)
}

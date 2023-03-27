package _092__反转链表_II

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	preHead := &ListNode{Next: head}
	pre := preHead
	cur := head
	for i := 1; i < left; i++ {
		pre = cur
		cur = cur.Next
	}
	//pre定位了外围链表的左侧

	var newTial *ListNode
	newHead := newTial
	//cur 已经是翻转的起始
	newTial = cur
	for i := 1; i <= right-left+1; i++ {
		next := cur.Next
		cur.Next = newHead
		newHead = cur
		cur = next
	}
	//cur指向外围链表的右侧
	pre.Next = newHead
	newTial.Next = cur
	return preHead.Next
}

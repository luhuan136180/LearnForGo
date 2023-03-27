package _203__移除链表元素

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {

	HPre := &ListNode{Next: head}
	cur := HPre
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else { //安全，后移
			cur = cur.Next
		}
	}
	return HPre.Next
}

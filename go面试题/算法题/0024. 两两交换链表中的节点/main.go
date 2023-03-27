package _024__两两交换链表中的节点

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	//head=list[i]
	//pre=list[i-1]
	pre := dummy
	for head != nil && head.Next != nil {
		pre.Next = head.Next
		temp := head.Next.Next
		head.Next.Next = head
		head.Next = temp
		pre = head
		head = temp
	}
	return dummy.Next

}

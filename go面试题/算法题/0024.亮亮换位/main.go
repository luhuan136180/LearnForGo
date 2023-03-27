package _024_亮亮换位

type ListNode struct {
	Val  int
	Next *ListNode
}

//递归法
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}

//迭代
func swapPairs2(head *ListNode) *ListNode {
	headPre := &ListNode{Next: head}
	pre := headPre
	cur := head
	for cur != nil && cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = cur
		pre.Next = next
		pre = cur
		cur = cur.Next

	}
	return headPre.Next
}

package main

func reverseKGroup2(head *ListNode, k int) *ListNode {
	preHead := &ListNode{Next: head}
	cur := head
	pre := preHead

	for cur != nil {
		start, end := cur, cur //设置翻转的起始和终结点
		for i := 1; i < k; i++ {
			end = end.Next

		}
		start, end = reverse2(start, end)
		pre.Next = start
		pre = end
	}

	return preHead.Next
}

func reverse2(start, end *ListNode) (Nstart, Nend *ListNode) {
	var pre *ListNode
	cur := start
	pre = end.Next
	for pre != end {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next

	}
	return end, start
}

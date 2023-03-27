package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next

	}
	return pre
}

func reverseList2(head *ListNode) *ListNode {
	return help(nil, head)
}

func help(pre, cur *ListNode) *ListNode {
	if cur == nil {
		return pre // 遍历结束
	}
	next := cur.Next
	cur.Next = pre
	return help(cur, next)
}

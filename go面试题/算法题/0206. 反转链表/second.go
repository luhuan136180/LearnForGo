package main

func reverseList3(head *ListNode) *ListNode {
	var preTial *ListNode //设置一个虚拟尾结点
	cur := head

	for cur != nil { //cur不是尾节点
		next := cur.Next
		cur.Next = preTial
		preTial = cur
		cur = next

	}
	return preTial
}

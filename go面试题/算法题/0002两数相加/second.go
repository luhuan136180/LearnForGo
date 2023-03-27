package main

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	//设置根节点
	root := &ListNode{
		Val: -1,
	}
	cur, carry := root, 0

	for l1 != nil || l2 != nil {
		var x, y int
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}

		res := x + y + carry

		cur.Next.Val = (res) % 10
		cur = cur.Next
		carry = (res) % 10
	}
	if carry != 0 {
		cur.Next.Val = carry
	}

	return root.Next
}

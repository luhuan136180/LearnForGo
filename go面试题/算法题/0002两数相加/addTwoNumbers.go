package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{
		Val: -1,
	}
	cur, carry := root, 0
	for l1 != nil || l2 != nil {
		x, y := 0, 0
		if l1 != nil {
			x = l1.Val
		}
		if l2 != nil {
			y = l2.Val
		}
		res := x + y + carry

		cur.Next = &ListNode{Val: res % 10}
		cur = cur.Next
		carry = res / 10

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

	}

	if carry != 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return root.Next
}

package _021__合并两个有序链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	temp1, temp2 := list1, list2
	headpre := &ListNode{Next: list1}
	pre := headpre
	for temp1 != nil || temp2 != nil {
		if temp1 == nil {
			pre.Next = temp2
			return headpre.Next
		}
		if temp2 == nil {
			pre.Next = temp1
			return headpre.Next
		}
		if temp1.Val <= temp2.Val {
			pre.Next = temp1
			next1 := temp1.Next
			temp1.Next = temp2
			pre = pre.Next
			temp1 = next1
		} else {
			pre.Next = temp2
			next2 := temp2.Next
			temp2.Next = temp1
			pre = pre.Next
			temp2 = next2
		}

	}
	return headpre.Next
}

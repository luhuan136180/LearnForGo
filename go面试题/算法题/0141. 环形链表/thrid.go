package main

func hasCycle4(head *ListNode) bool {
	fast, slow := head, head
	//有且只有一个节点的话无法构成环
	if head == nil || head.Next == nil {
		return false
	}
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//常规
func hasCycle(head *ListNode) bool {
	arr := make(map[*ListNode]int)

	cur := head
	for cur != nil {
		if _, ok := arr[cur]; !ok {
			arr[cur] = 1
			cur = cur.Next
		} else {
			return true
		}

	}
	return false
}

//快慢指针
func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next

	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return true
}

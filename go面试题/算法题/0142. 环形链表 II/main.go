package _142__环形链表_II

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	var cur *ListNode
	cur = head
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			for slow != cur {
				slow = slow.Next
				cur = cur.Next
			}
			return cur
		}
	}
	return nil
}

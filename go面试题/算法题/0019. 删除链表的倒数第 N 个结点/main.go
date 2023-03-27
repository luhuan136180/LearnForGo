package _019__删除链表的倒数第_N_个结点

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	preHead := &ListNode{Next: head}
	slow := head
	fast := head
	for i := 1; i < n; i++ {
		fast = fast.Next
	}
	pre := preHead

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
		pre = pre.Next
	}
	pre.Next = slow.Next
	slow.Next = nil
	return preHead.Next
}

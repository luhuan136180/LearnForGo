package main

func main() {

}

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//线性表发
func reorderList(head *ListNode) {
	arr := []*ListNode{}
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}
	i, j := 0, len(arr)-1
	for i < j {
		arr[i].Next = arr[j]
		i++
		if i == j {
			break
		}
		arr[j].Next = arr[i]
		j--
	}
	arr[i].Next = nil
}

//寻找链表中点 + 链表逆序 + 合并链表
func reorderList2(head *ListNode) {
	if head == nil {
		return
	}
	//寻找中间节点
	mid := middleNode(head)
	l1 := head
	//将后半段的起点指针调整好
	l2 := mid.Next
	mid.Next = nil
	//以中间节点为基准对后半段进行链表翻转
	l2 = reverseList(l2)
	//将前，后半段交叉整合为一条链
	mergeList(l1, l2)
}
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow //slow 标注为上半段的最后一个元素（偶数），正中间（奇数）
}
func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func mergeList(l1, l2 *ListNode) {
	var l1tmp, l2tmp *ListNode
	for l1 != nil && l2 != nil {
		l1tmp = l1.Next
		l2tmp = l2.Next

		l1.Next = l2
		l1 = l1tmp

		l2.Next = l1
		l2 = l2tmp

	}
}

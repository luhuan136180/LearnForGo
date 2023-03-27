package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	cur := head                   //设置移动节点
	hair := &ListNode{Next: head} //设置空节点指向head
	pre := hair

	for cur != nil {
		start := cur
		temp := pre //设置移动节点的前置节点
		for i := 0; i < k; i++ {
			temp = cur
			cur = cur.Next
			if cur == nil && i != k-1 {
				return hair.Next
			}
		}
		//截取到移动节点的前置节点的k的节点做反转，

		start, temp = reverse(start, temp)
		pre.Next = start
		pre = temp
	}
	return hair.Next
}

func reverse(start, end *ListNode) (Nstart, Nend *ListNode) {
	cur := start
	pre := end.Next
	for pre != end {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return end, start
}

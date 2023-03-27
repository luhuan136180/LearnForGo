package _707__设计链表

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
type ListNode struct {
	Val  int
	Next *ListNode
}
type MyLinkedList struct {
	head *ListNode //虚头结点
	size int
}

func Constructor() MyLinkedList {
	preHead := &ListNode{}
	return MyLinkedList{head: preHead, size: 0}
}

func (this *MyLinkedList) Get(index int) int {
	if this.size == 0 || index > this.size-1 || index < 0 {
		return -1
	}
	//在范围内
	cur := this.head //虚头结点

	for i := 0; i <= index; i++ { //i=0头结点
		cur = cur.Next
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.size++
	newHead := &ListNode{Val: val}
	newHead.Next = this.head.Next
	this.head.Next = newHead
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.size++
	cur := this.head.Next
	for cur.Next != nil {
		cur = cur.Next
	}
	newTail := &ListNode{Val: val}
	cur.Next = newTail
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == this.size {
		this.AddAtTail(val)
	}
	if index > this.size {
		return
	}
	new := &ListNode{Val: val}
	cur := this.head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	new.Next = cur.Next
	cur.Next = new
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this.size == 0 || index > this.size-1 || index < 0 {
		return
	}
	cur := this.head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	pre := cur.Next
	cur.Next = cur.Next.Next
	pre.Next = nil
}

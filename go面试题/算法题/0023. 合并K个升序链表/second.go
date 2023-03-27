package _023__合并K个升序链表

import "container/heap"

type ListNodeHeap []*ListNode

func (h ListNodeHeap) Len() int           { return len(h) }
func (h ListNodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h ListNodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *ListNodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}
func (h *ListNodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists2(lists []*ListNode) *ListNode {
	var h = new(ListNodeHeap)
	for _, list := range lists {
		if list != nil {
			heap.Push(h, list)
		}
	}
	var dummy = new(ListNode)
	var head = dummy
	for h.Len() > 0 {
		top := heap.Pop(h).(*ListNode)
		if top.Next != nil {
			heap.Push(h, top.Next)
		}
		head.Next, head = top, top
	}
	return dummy.Next
}

package _116__填充每个节点的下一个右侧节点指针

import "container/list"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	queue := list.New()
	if root == nil {
		return root
	}
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		temparr := []*Node{}
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*Node)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			temparr = append(temparr, node)
		}
		if len(temparr) > 1 {
			for j := 0; j < len(temparr)-1; j++ {
				temparr[j].Next = temparr[j+1]
			}
		}
	}
	return root
}

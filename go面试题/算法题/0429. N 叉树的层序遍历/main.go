package _429__N_叉树的层序遍历

import "container/list"

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	queue := list.New()
	var ans [][]int
	if root == nil {
		return ans
	}
	queue.PushBack(root)
	for queue.Len() > 0 {
		lenth := queue.Len()
		var tmp []int
		for t := 0; t < lenth; t++ {
			myNode := queue.Remove(queue.Front()).(*Node)
			tmp = append(tmp, myNode.Val)
			for i := 0; i < len(myNode.Children); i++ {
				queue.PushBack(myNode.Children[i])
			}
		}
		ans = append(ans, tmp)
	}
	return ans
}

package _104__二叉树的最大深度

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	queue := list.New()
	depth := 0
	if root == nil {
		return depth
	}
	queue.PushBack(root)
	for queue.Len() > 0 {
		depth++ //当前层有元素，进入循环，所以深度+1
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}

		}

	}
	return depth
}

package _111__二叉树的最小深度

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	queue := list.New()
	depth := 0
	if root == nil {
		return 0
	}
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		depth++
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)

			if node.Left == nil && node.Right == nil {
				return depth
			}
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

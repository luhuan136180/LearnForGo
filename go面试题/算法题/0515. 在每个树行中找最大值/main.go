package _515__在每个树行中找最大值

import (
	"container/list"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	queue := list.New()
	var res []int

	if root == nil {
		return res
	}
	queue.PushBack(root)

	for queue.Len() > 0 {
		length := queue.Len()
		max := math.MinInt64
		for i := 0; i < length; i++ {

			node := queue.Remove(queue.Front()).(*TreeNode)
			if max < node.Val {
				max = node.Val
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		res = append(res, max)
	}
	return res
}

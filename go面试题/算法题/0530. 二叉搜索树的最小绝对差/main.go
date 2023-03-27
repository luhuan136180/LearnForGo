package _530__二叉搜索树的最小绝对差

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	min := math.MaxInt64
	var pre *TreeNode

	var tralve func(node *TreeNode)

	tralve = func(node *TreeNode) {
		if node == nil {
			return
		}
		tralve(node.Left)
		if pre != nil && node.Val-pre.Val < min {
			min = node.Val - pre.Val
		}
		pre = node
		tralve(node.Right)
	}

	tralve(root)
	return min
}

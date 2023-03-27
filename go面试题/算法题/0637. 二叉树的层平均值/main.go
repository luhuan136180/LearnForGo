package _637__二叉树的层平均值

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	var ans []float64
	arr := list.New()
	arr.PushBack(root)
	for arr.Len() > 0 {
		lenth := arr.Len()
		var sum float64
		for i := 0; i < lenth; i++ {
			node := arr.Remove(arr.Front()).(*TreeNode)
			if node.Left != nil {
				arr.PushBack(node.Left)
			}
			if node.Right != nil {
				arr.PushBack(node.Right)
			}
			sum += float64(node.Val)
		}
		ans = append(ans, float64(sum/float64(lenth)))
	}

	return ans
}

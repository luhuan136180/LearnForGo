package _501__二叉搜索树中的众数

import (
	"container/list"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//不是二叉搜索树的做法
func findMode(root *TreeNode) []int {
	ans := make(map[int]int, 0)
	var res []int
	queue := list.New()
	queue.PushBack(root)
	max := math.MinInt64
	if root == nil {
		return res
	}
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			ans[node.Val]++
			if ans[node.Val] > max {
				max = ans[node.Val]
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	for i, val := range ans {
		if val == max {
			res = append(res, i)
		}
	}

	return res
}

//二叉搜索树——使用中序遍历推导出有序数组
func findMode2(root *TreeNode) []int {
	res := make([]int, 0)
	max := math.MinInt64
	count := 1
	var pre *TreeNode
	var travle func(node *TreeNode) //递归函数
	travle = func(node *TreeNode) {
		if node == nil {
			return
		}
		travle(node.Left)

		if pre != nil && pre.Val == node.Val {
			//有前节点，且val相等
			count++
		} else {
			count = 1 //重置
		}
		if count >= max {
			if count > max && len(res) > 0 {
				res = []int{node.Val} //数组重置

			} else {
				res = append(res, node.Val)
			}
			max = count
		}
		pre = node

		travle(node.Right)
	}

	travle(root)
	return res
}

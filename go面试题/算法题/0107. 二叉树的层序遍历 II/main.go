package _107__二叉树的层序遍历_II

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	var curlist []*TreeNode
	curlist = append(curlist, root)
	for len(curlist) > 0 {
		var nextlist []*TreeNode
		var vals []int //当前层的数组
		for _, node := range curlist {
			vals = append(vals, node.Val)
			if node.Left != nil {
				nextlist = append(nextlist, node.Left)
			}
			if node.Right != nil {
				nextlist = append(nextlist, node.Right)
			}
		}
		ans = append(ans, vals)
		curlist = nextlist
	}
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-i-1] = ans[len(ans)-i-1], ans[i]
	}
	return ans
}

func reserve(res [][]int) [][]int {
	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
	}
	return res
}

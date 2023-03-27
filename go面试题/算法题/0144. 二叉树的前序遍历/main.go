package _144__二叉树的前序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var ans []int

	var tranversal func(cur *TreeNode)
	tranversal = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		ans = append(ans, cur.Val)
		tranversal(cur.Left)
		tranversal(cur.Right)
	}
	tranversal(root)
	return ans
}

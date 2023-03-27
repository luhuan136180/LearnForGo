package _222__完全二叉树的节点个数

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftnum := countNodes(root.Left)
	rightnum := countNodes(root.Right)
	treenum := 1 + leftnum + rightnum
	return treenum
}

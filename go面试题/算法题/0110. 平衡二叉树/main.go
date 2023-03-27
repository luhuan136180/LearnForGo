package _110__平衡二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if getHeight(root) == -1 {
		return false
	}
	return true
}

func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftHeight := getHeight(node.Left)
	rightHeight := getHeight(node.Right)

	if leftHeight == -1 || rightHeight == -1 {
		return -1
	}
	if leftHeight-rightHeight > 1 || rightHeight-leftHeight > 1 {
		return -1
	}
	return 1 + max(leftHeight, rightHeight)
}

func max(l, r int) int {
	if l > r {
		return l
	}
	return r
}

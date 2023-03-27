package _404__左叶子之和

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//中序
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftValue := sumOfLeftLeaves(root.Left)

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		leftValue = leftValue + root.Left.Val
	}
	rightValue := sumOfLeftLeaves(root.Right)
	return leftValue + rightValue
}

//

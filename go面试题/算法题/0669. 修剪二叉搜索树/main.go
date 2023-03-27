package _669__修剪二叉搜索树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//如果root（当前节点）的元素小于low的数值，那么应该递归右子树，并返回右子树符合条件的头结点。
//如果root(当前节点)的元素大于high的，那么应该递归左子树，并返回左子树符合条件的头结点。
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		right := trimBST(root.Right, low, high)
		return right
	}
	if root.Val > high {
		left := trimBST(root.Left, low, high)
		return left
	}

	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)

	return root

}

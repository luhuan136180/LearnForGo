package _701__二叉搜索树中的插入操作

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {

	var travle func(node *TreeNode, val int) *TreeNode

	travle = func(node *TreeNode, val int) *TreeNode {

		if node == nil {
			return &TreeNode{Val: val}
		}
		if node.Val > val {
			node.Left = travle(node.Left, val)
		} else {
			node.Right = travle(node.Right, val)
		}
		return node

	}

	root = travle(root, val)
	return root
}

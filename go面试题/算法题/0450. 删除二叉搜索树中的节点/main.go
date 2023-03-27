package _450__删除二叉搜索树中的节点

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//第一种情况：没找到删除的节点，遍历到空节点直接返回了 找到删除的节点
//第二种情况：左右孩子都为空（叶子节点），直接删除节点， 返回NULL为根节点
//第三种情况：删除节点的左孩子为空，右孩子不为空，删除节点，右孩子补位，返回右孩子为根节点
//第四种情况：删除节点的右孩子为空，左孩子不为空，删除节点，左孩子补位，返回左孩子为根节点
//第五种情况：左右孩子节点都不为空，则将删除节点的左子树头结点（左孩子）放到删除节点的右子树的最左面节点的左孩子上，返回删除节点右孩子为新的根节点。
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == key {
		if root.Left == nil {
			//第二，三种
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			cur := root.Right
			for cur.Left != nil {
				cur = cur.Left
			}
			cur.Left = root.Left // 把要删除的节点（root）左子树放在cur的左孩子的位置
			root = root.Right

		}

	}
	root.Left = deleteNode(root.Left, key)
	root.Right = deleteNode(root.Right, key)

	return root
}

func deleteNode2(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	}
	//key==root.Val
	if root.Right == nil {
		return root.Left
	}
	if root.Left == nil {
		return root.Right
	}
	cur := root.Right
	if cur.Left != nil {
		cur = cur.Left
	}
	root.Val = cur.Val
	root.Right = deleteNodeNext(root.Right)
	return root
}
func deleteNodeNext(root *TreeNode) *TreeNode {
	if root.Left == nil {
		pRight := root.Right
		root.Right = nil
		return pRight
	}
	root.Left = deleteNodeNext(root.Left)
	return root
}

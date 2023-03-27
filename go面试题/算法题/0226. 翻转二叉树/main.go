package _226__翻转二叉树

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归版本的前序遍历
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left

	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

//递归版本的后序遍历
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	invertTree2(root.Left)
	invertTree2(root.Right)
	root.Left, root.Right = root.Right, root.Left

	return root
}

//层序遍历
func invertTree3(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	queue := list.New()
	node := root
	queue.PushBack(node)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			e := queue.Remove(queue.Front()).(*TreeNode)
			e.Left, e.Right = e.Right, e.Left
			if e.Left != nil {
				queue.PushBack(e.Left)
			}
			if e.Right != nil {
				queue.PushBack(e.Right)
			}
		}
	}
	return root
}

//func invertTree4(root *TreeNode) *TreeNode {
//
//}

package _98__验证二叉搜索树

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	list := make([]int, 0)
	var travle func(node *TreeNode, list []int) []int

	travle = func(node *TreeNode, list []int) []int {
		if node == nil {
			return list
		}
		list = travle(node.Left, list)
		list = append(list, node.Val)
		list = travle(node.Right, list)
		return list
	}
	list = travle(root, list)
	fmt.Println(list)
	for i := 1; i < len(list); i++ {
		if list[i-1] >= list[i] {
			return false
		}
	}
	return true
}

package _145__二叉树的后序遍历

//迭代法
func postorderTraversal2(root *TreeNode) []int {
	var stack []*TreeNode
	var ans []int
	if root == nil {
		return ans
	}
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		ans = append(ans, node.Val)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	ans = reverse(ans)
	return ans
}
func reverse(s []int) []int {
	l, r := 0, len(s)-1

	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
	return s
}

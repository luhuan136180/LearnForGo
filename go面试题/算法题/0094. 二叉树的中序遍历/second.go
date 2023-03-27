package _094__二叉树的中序遍历

//迭代法

func inorderTraversal2(root *TreeNode) []int {

	var stack []*TreeNode
	var ans []int
	//需要用到指针
	cur := root
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]    //栈顶元素，表示该元素左子节点没有元素
			stack = stack[:len(stack)-1] //弹出
			ans = append(ans, cur.Val)
			cur = cur.Right
		}
	}
	return ans
}

//左中右——压入栈的顺序则为右中左
func inorderTraversal3(root *TreeNode) []int {
	var stack []*TreeNode
	var res []int
	if root == nil {
		return res
	}
	stack = append(stack, root)
	if len(stack) > 0 {
		node := stack[len(stack)-1] //最后一个元素
		stack = stack[:len(stack)-1]
		if node.Val == 0 {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, node.Val)

		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		stack = append(stack, node)
		stack = append(stack, &TreeNode{})
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return res
}

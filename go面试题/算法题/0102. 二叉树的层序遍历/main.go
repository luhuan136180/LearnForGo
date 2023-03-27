package _102__二叉树的层序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归法
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	depth := 0

	var order func(root *TreeNode, depth int)

	order = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if len(res) == depth {
			res = append(res, []int{})
		}
		res[depth] = append(res[depth], root.Val)

		order(root.Left, depth+1)
		order(root.Right, depth+1)
	}

	order(root, depth)

	return res
}

/**
102. 二叉树的层序遍历：使用切片模拟队列，易理解
*/
func levelOrder2(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var curlist []*TreeNode //当前层——root层
	curlist = append(curlist, root)
	for len(curlist) > 0 {
		nextlist := []*TreeNode{} // 准备通过当前层生成下一层
		vals := []int{}

		for _, node := range curlist {
			vals = append(vals, node.Val) // 收集当前层的值
			// 收集下一层的节点
			if node.Left != nil {
				nextlist = append(nextlist, node.Left)
			}
			if node.Right != nil {
				nextlist = append(nextlist, node.Right)
			}

		}
		res = append(res, vals) //一层遍历结束
		curlist = nextlist      // 将下一层变成当前层
	}
	return res
}

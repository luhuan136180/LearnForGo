package _144__二叉树的前序遍历

import (
	"container/list"
)

//迭代法
//迭代器与要跟迭代元素直接相关，构建迭代用的栈时需要判断类型，不能盲目做
func preorderTraversal2(root *TreeNode) []int {
	var stack []*TreeNode
	var ans []int

	if root == nil {
		return ans
	}
	stack = append(stack, root)

	for len(stack) > 0 {
		node := stack[len(stack)-1]  //取租后一个元素
		stack = stack[:len(stack)-1] //弹出

		ans = append(ans, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return ans

}

//统一迭代法
func preorderTraversal3(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var stack = list.New() //创建栈
	res := []int{}

	stack.PushBack(root)
	var node *TreeNode
	for stack.Len() > 0 {
		e := stack.Back()
		stack.Remove(e)     //弹出元素
		if e.Value == nil { // 如果为空，则表明是需要处理中间节点
			e = stack.Back() //弹出元素（即中间节点）
			stack.Remove(e)  //删除中间节点
			node = e.Value.(*TreeNode)
			res = append(res, node.Val) //将中间节点加入到结果集中
			continue                    //继续弹出栈中下一个节点
		}
		node = e.Value.(*TreeNode)
		//压栈顺序：右左中
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
		stack.PushBack(node) //中间节点压栈后再压入nil作为中间节点的标志符
		stack.PushBack(nil)
	}
	return res
}

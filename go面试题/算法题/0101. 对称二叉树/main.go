package _101__对称二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归法
//确定递归函数的参数和返回值:bool compare(TreeNode* left, TreeNode* right)

//确定终止条件
//节点为空的情况有：（注意我们比较的其实不是左孩子和右孩子，所以如下我称之为左节点右节点）
//
//左节点为空，右节点不为空，不对称，return false
//左不为空，右为空，不对称 return false
//左右都为空，对称，返回true

//确定单层递归的逻辑
//此时才进入单层递归的逻辑，单层递归的逻辑就是处理 左右节点都不为空，且数值相同的情况。
//
//比较二叉树外侧是否对称：传入的是左节点的左孩子，右节点的右孩子。
//比较内测是否对称，传入左节点的右孩子，右节点的左孩子。
//如果左右都对称就返回true ，有一侧不对称就返回false 。
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compare(root.Left, root.Right)

}
func compare(left, right *TreeNode) bool {
	if left != nil && right == nil {
		return false
	} else if left == nil && right != nil {
		return false
	} else if left == nil && right == nil {
		return true
	} else if left.Val != right.Val {
		return false
	}

	//一下为左右节点都不为空，且相等
	//递归遍历左右节点的子节点比较
	outside := compare(left.Left, right.Right)
	inside := compare(left.Right, right.Left)
	return outside && inside
}

//迭代法
func isSymmetric2(root *TreeNode) bool {
	//切片模拟队列
	var queue []*TreeNode
	if root != nil {
		queue = append(queue, root.Left, root.Right)
	}
	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		queue = append(queue, left.Left, right.Right, right.Left, left.Right)
	}
	return true
}

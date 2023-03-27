package _112__路径总和

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归，回溯初版
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	var travle func(node *TreeNode, ans int)
	var res []int //所有路径值的集合
	travle = func(node *TreeNode, ans int) {
		if node.Right == nil && node.Left == nil {
			//找到一个子节点
			ans = ans + node.Val //得到路径值
			res = append(res, ans)
			return //回溯
		}
		ans = ans + node.Val
		if node.Left != nil {
			travle(node.Left, ans)
		}
		if node.Right != nil {
			travle(node.Right, ans)
		}

	}
	travle(root, 0)
	for _, val := range res {
		if val == targetSum {
			return true
		}
	}
	return false
}

//升级版递归
func hasPathSum2(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	var travle func(node *TreeNode, ans int)
	res := make(map[int]bool) //所有路径值的集合
	travle = func(node *TreeNode, ans int) {
		ans = ans + node.Val //得到当前路径值
		if node.Right == nil && node.Left == nil {
			//找到一个子节点
			res[ans] = true
			return //回溯
		}
		if node.Left != nil {
			travle(node.Left, ans)
		}
		if node.Right != nil {
			travle(node.Right, ans)
		}

	}
	travle(root, 0)

	return res[targetSum]
}

//递归
func hasPathSum3(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	// 将targetSum在遍历每层的时候都减去本层节点的值
	targetSum -= root.Val
	if root.Left == nil && root.Right == nil && targetSum == 0 {
		// 如果剩余的targetSum为0, 则正好就是符合的结果
		return true
	}
	// 否则递归找
	return hasPathSum3(root.Left, targetSum) || hasPathSum3(root.Right, targetSum)
}

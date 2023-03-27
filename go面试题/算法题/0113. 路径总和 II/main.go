package _113__路径总和_II

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	var travle func(node *TreeNode, ans int, path []int)
	var path []int

	travle = func(node *TreeNode, ans int, path []int) {
		ans += node.Val
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil {
			if ans == targetSum {
				copy := make([]int, 0)
				for _, val := range path {
					copy = append(copy, val)
				}
				res = append(res, copy)
			}
			//return      //可要可不要
		}
		if node.Left != nil {
			travle(node.Left, ans, path)
		}
		if node.Right != nil {
			travle(node.Right, ans, path)
		}
		//path = path[:len(path)-1]
	}

	travle(root, 0, path)
	return res
}

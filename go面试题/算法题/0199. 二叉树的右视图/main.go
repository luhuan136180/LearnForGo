package _199__二叉树的右视图

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//用一层一列表的思路就是遍历到每层最后是输入到结果集，len()控制
//用一个数组模拟一维列表 并逻辑建立二维列表
/*********************************************************/
//每次进入循环时，先取得len()，遍历，i=len()-1时，当层结束，后面的为下一层元素

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	list := list.New()
	list.PushBack(root)

	for list.Len() > 0 {
		lenth := list.Len() //取得当前层的元素个数，标记

		for i := 0; i < lenth; i++ {
			//返回的是一个any，需要断言
			node := list.Remove(list.Front()).(*TreeNode)
			if node.Left != nil {
				list.PushBack(node.Left)
			}
			if node.Right != nil {
				list.PushBack(node.Right)
			}

			// // 取每层的最后一个元素，添加到结果集中
			if i == lenth-1 {
				res = append(res, node.Val)
			}
		}
	}
	return res
}

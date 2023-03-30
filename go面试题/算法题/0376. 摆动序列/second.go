package main

func wiggleMaxLength3(nums []int) int {
	n := len(nums)
	if n < 2 { //一个节点时
		return n
	}
	//两个节点时，如果平行也不添加子序列长度
	ans := 1 //起始为一
	pre := nums[1] - nums[0]
	if pre != 0 {
		return 2
	}

	for i := 2; i < n; i++ {
		cur := nums[i] - nums[i-1]
		if cur > 0 && pre < 0 || cur < 0 && pre > 0 {
			ans++
		}
		pre = cur
	}

	return ans
}

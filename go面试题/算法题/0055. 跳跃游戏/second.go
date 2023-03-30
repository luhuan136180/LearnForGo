package main

func canJump2(nums []int) bool {
	cover := 0 //可覆盖范围

	for i := 0; i <= cover; i++ {
		if i+nums[i] > cover {
			cover = i + nums[i]
		}
		if cover >= len(nums)-1 {
			return true
		}
	}
	return false
}

//动态规划
func canJump3(nums []int) bool {
	//dp[]表示在index的最远可覆盖距离
	dp := make([]int, len(nums))
	//3初始化
	dp[0] = nums[0]

	//2.确定递推公式 max(dp[])
	for i := 1; i < len(nums); i++ {

		if dp[i-1] >= i {
			dp[i] = max(dp[i-1], i+nums[i])
		}
		if dp[i] >= len(nums)-1 {
			return true
		}

	}
	return false
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

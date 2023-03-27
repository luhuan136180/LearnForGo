package _198__打家劫舍

func rob2(nums []int) int {
	//
	//dp[j] = max(dp[j-1],dp[j-2]+nums[i]
	//初始化：dp[0] =
	lenth := len(nums)
	dp := make([]int, lenth)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < lenth; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[lenth-1]
}

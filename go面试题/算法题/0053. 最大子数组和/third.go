package main

//贪心法
func maxSubArray3(nums []int) int {
	res := nums[0]
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > res {
			res = sum
		}
		if sum <= 0 {
			sum = 0
		}
	}
	return res
}

//动态规划
func maxSubArray4(nums []int) int {
	//dp数组
	//	公式
	//	初始化
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

package main

func canJump(nums []int) bool {
	if len(nums) <= 0 {
		return true
	}
	dp := make([]bool, len(nums))
	dp[0] = true
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && nums[j]+j >= i {
				dp[i] = true
				break //退出
			}
		}
	}
	return dp[len(nums)-1]
}

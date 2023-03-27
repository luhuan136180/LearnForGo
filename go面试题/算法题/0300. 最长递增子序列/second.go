package _300__最长递增子序列

func lengthOfLIS2(nums []int) int {
	dp := make([]int, len(nums))

	dp[0] = 1

	ans := 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[i-j]+1)
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}

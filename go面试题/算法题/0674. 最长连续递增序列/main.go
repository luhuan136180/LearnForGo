package _674__最长连续递增序列

func findLengthOfLCIS(nums []int) int {
	lenth := len(nums)
	dp := make([]int, lenth)

	for i, _ := range dp {
		dp[i] = 1
	}
	ans := 1
	for i := 0; i < lenth-1; i++ {
		if nums[i+1] > nums[i] {
			dp[i+1] = dp[i] + 1
		}
		if dp[i+1] > ans {
			ans = dp[i+1]
		}
	}
	return ans
}

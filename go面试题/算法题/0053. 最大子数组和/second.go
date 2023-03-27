package main

func maxSubArray2(nums []int) int {
	ans := 0
	dp := make([]int, len(nums))

	dp[0] = nums[0]

	for i := 0; i < len(nums); i++ {
		//递推公式表示：延续之前的数组和从自己开始的新数组，取最大值
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		if ans < dp[i] {
			ans = dp[i]
		}
	}
	return ans

}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

package main

func rob2(nums []int) int {
	lenth := len(nums)

	if lenth == 0 {
		return 0
	}
	if lenth == 1 {
		return nums[0]
	}
	//分两种情况，考虑队首，不包含队尾
	//考虑队尾，不包含队首
	first := maxMoney(nums, 0, lenth-2)
	second := maxMoney(nums, 1, lenth-1)
	ans := max(first, second)
	return ans

}

func maxMoney(nums []int, start, end int) int {
	if start == end {
		return nums[start]
	}
	dp := make([]int, end-start+1)

	dp[0] = nums[start]
	dp[1] = max(dp[0], nums[start+1])
	for i := 2; i < len(dp); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[start+i])
	}

	return dp[end-start]
}

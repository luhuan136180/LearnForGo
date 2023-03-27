package _309__最佳买卖股票时机含冷冻期

//四个状态：1.今天持有 2.今天保持卖出状态，且不是冷冻期，至少前天就卖了；3.今天卖出股票 4.今天为冷冻期
func maxProfit(prices []int) int {
	lenth := len(prices)
	dp := make([][]int, lenth)

	for i, _ := range dp {
		dp[i] = make([]int, 4)
	}
	dp[0][0] = -prices[0]
	for i := 1; i < lenth; i++ {
		dp[i][0] = max(dp[i-1][0], max(dp[i-1][1]-prices[i], dp[i-1][3]-prices[i]))
		dp[i][1] = max(dp[i-1][1], dp[i-1][3])
		dp[i][2] = dp[i-1][0] + prices[i]
		dp[i][3] = dp[i-1][2]
	}

	return max(dp[lenth-1][3], max(dp[lenth-1][1], dp[lenth-1][2]))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

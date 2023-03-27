package _188__买卖股票的最佳时机_IV

func maxProfit(k int, prices []int) int {
	lenth := len(prices)
	dp := make([][]int, lenth)
	for i, _ := range dp {
		dp[i] = make([]int, k*2+1)
	}

	for i := 1; i < len(dp[0]); i += 2 {
		dp[0][i] = -prices[0]
	}

	for i := 1; i < lenth; i++ {
		for j := 0; j < 2*k-1; j += 2 {
			dp[i][j+1] = max(dp[i-1][j]-prices[i], dp[i-1][j+1])
			dp[i][j+2] = max(dp[i-1][j+1]+prices[i], dp[i-1][j+2])

		}
	}
	return dp[lenth-1][2*k]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

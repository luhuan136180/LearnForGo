package _714__买卖股票的最佳时机含手续费

func maxProfit(prices []int, fee int) int {
	dp := make([][]int, len(prices))
	for i, _ := range dp {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = -prices[0] - fee
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i]-fee)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}

	return dp[len(prices)-1][1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

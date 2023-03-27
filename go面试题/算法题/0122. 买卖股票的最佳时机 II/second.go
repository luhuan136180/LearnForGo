package main

func maxProfit2(prices []int) int {
	lenth := len(prices)
	if lenth == 0 {
		return 0
	}
	dp := make([][]int, lenth)
	for i, _ := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < lenth; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}

	return dp[lenth-1][1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

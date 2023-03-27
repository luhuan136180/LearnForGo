package main

func maxProfit2(prices []int) int {
	lenth := len(prices)
	if lenth == 0 {
		return 0
	}
	dp := make([][]int, lenth)
	for i, _ := range dp {
		dp[i] = make([]int, 5)
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[0][3] = -prices[0]

	for i := 1; i < lenth; i++ {
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
		dp[i][2] = max(dp[i-1][1]+prices[i], dp[i-1][2])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}
	return dp[lenth-1][4]
}

//
//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}

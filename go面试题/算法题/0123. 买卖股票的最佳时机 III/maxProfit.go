package main

func maxProfit(prices []int) int {

	dp := make([][]int, len(prices)) //创建一个dp；值表示当前状态下的手持金额
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 5)

	}

	//初始化
	//dp【1】【j】：0,-1,0,-1,0
	dp[0][1] = -prices[0]
	dp[0][3] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])

	}

	return dp[len(prices)-1][4]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

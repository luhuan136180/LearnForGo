package _279__完全平方数

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	//找最少数量，初始化为最大值
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt64
	}

	for i := 1; i*i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j >= i*i {
				dp[j] = min(dp[j], dp[j-i*i]+1)
			}
		}
	}

	return dp[n]

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

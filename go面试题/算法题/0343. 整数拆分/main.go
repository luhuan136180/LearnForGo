package _343__整数拆分

func integerBreak(n int) int {
	//1dp[i]=表示 n=i时的最大乘积
	//2。dp[i] = max(j*(i-j),j*dp[i-j])     (j<i)
	//3初始化:dp[0]=0  dp[1] = 0  dp[2] = 1
	//4.遍历顺序:dp[i] 是依靠 dp[i - j]的状态，所以遍历i⼀定是从前向后遍历;j从1开始
	//5.
	if n == 0 || n == 1 {
		return 0
	}
	dp := make([]int, n+1)
	dp[2] = 1

	for i := 3; i < len(dp); i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max(dp[i-j]*j, j*(i-j)))
		}
	}

	return dp[n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

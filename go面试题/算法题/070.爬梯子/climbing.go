package main

//斐波那契，递归求解
func climbStairs(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	return climbStairs(n-1) + climbStairs(n-2)
}
func climbStairs2(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

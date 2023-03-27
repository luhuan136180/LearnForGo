package _392__判断子序列

func isSubsequence(s string, t string) bool {
	dp := make([][]int, len(s)+1)

	for i, _ := range dp {
		dp[i] = make([]int, len(t)+1)
	}
	ans := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			if s[i] == t[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = dp[i+1][j]
			}

			if ans < dp[i+1][j+1] {
				ans = dp[i+1][j+1]
			}
		}
	}
	return ans == len(s)
}

package _115__不同的子序列

func numDistinct(s string, t string) int {
	dp := make([][]int, len(s)+1)

	for i, _ := range dp {
		dp[i] = make([]int, len(t)+1)
	}

	for i := 0; i < len(dp); i++ {
		dp[i][0] = 1
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			if s[i] == t[j] {
				dp[i+1][j+1] = dp[i][j] + dp[i][j+1]
			} else {
				dp[i+1][j+1] = dp[i][j+1]
			}

		}
	}

	return dp[len(s)][len(t)]
}

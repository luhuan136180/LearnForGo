package _583__两个字符串的删除操作

func minDistance(word1 string, word2 string) int {
	//dp表示i，j时的最小步数
	dp := make([][]int, len(word1)+1)

	for i, _ := range dp {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 0; i < len(dp); i++ {
		dp[i][0] = i
	}
	for j := 0; j < len(dp[1]); j++ {
		dp[0][j] = j
	}

	for i := 0; i < len(word1); i++ {
		for j := 0; j < len(word2); j++ {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = min(dp[i][j+1]+1, min(dp[i+1][j]+1, dp[i][j]+2))

			}

		}
	}

	return dp[len(word1)][len(word2)]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

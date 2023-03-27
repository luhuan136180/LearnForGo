package _072_编辑距离

func minDistance2(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)

	for i, _ := range dp {
		dp[i] = make([]int, len(word2)+1)
	}

	for i := 0; i < len(dp); i++ {
		dp[i][0] = i
	}
	for j := 0; j < len(dp[0]); j++ {
		dp[0][j] = j
	}

	//
	for i := 0; i < len(word1); i++ {
		for j := 0; j < len(word2); j++ {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = Min(dp[i][j+1], Min(dp[i][j], dp[i+1][j])) + 1
			}
		}
	}

	return dp[len(word1)][len(word2)]
}
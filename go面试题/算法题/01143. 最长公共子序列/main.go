package _1143__最长公共子序列

//不是很懂
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i, _ := range dp {
		dp[i] = make([]int, len(text2)+1)
	}
	ans := 0
	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
			if ans < dp[i+1][j+1] {
				ans = dp[i+1][j+1]
			}

		}
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y

}

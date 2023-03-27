package _139__单词拆分

func wordBreak(s string, wordDict []string) bool {
	// 装满背包s的前几位字符的方式有几种
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 0; i <= len(s); i++ { // 背包
		for j := 0; j < len(wordDict); j++ { // 物品
			if i >= len(wordDict[j]) && wordDict[j] == s[i-len(wordDict[j]):i] {
				dp[i] += dp[i-len(wordDict[j])]
			}
		}
	}

	return dp[len(s)] > 0
}

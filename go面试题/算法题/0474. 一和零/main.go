package _474__一和零

func findMaxForm(strs []string, m int, n int) int {
	//定义数组
	dp := make([][]int, m+1)

	for i, _ := range dp {
		dp[i] = make([]int, n)
	}
	//遍历
	for i := 0; i < len(strs); i++ {
		zeroNum, oneNum := 0, 0
		//计算0,1的个数
		for _, val := range strs[i] {
			if val == '0' {
				zeroNum++
			}
		}
		oneNum = len(strs[i]) - zeroNum //求当前字符的0,1个数
		//从后往前遍历背包容量
		for j := m; j > zeroNum; j-- {
			for k := n; k > oneNum; k-- {
				dp[j][k] = max(dp[j][k], dp[j-zeroNum][k-oneNum])
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

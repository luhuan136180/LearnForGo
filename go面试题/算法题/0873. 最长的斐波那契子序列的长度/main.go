package _873__最长的斐波那契子序列的长度

func lenLongestFibSubseq(arr []int) int {
	//dp表示index位置的子序长度
	//

	//常规递归   需要复做
	n := len(arr)
	idxMap := make(map[int]int, 0)
	dp := make([][]int, n)
	for i, val := range arr {
		idxMap[val] = i
		dp[i] = make([]int, n)
	}
	var ans int
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			dp[i][j] = max(dp[i][j], 2)
			if k, ok := idxMap[arr[i]+arr[j]]; ok {
				dp[j][k] = dp[i][j] + 1
				ans = max(ans, dp[j][k])
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
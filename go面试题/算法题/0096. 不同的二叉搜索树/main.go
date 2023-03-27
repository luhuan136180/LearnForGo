package _096__不同的二叉搜索树

func numTrees(n int) int {
	//dp数组表示i个节点时的二叉搜索树总数
	//递推公式dp[i] += dp[j - 1] * dp[i - j]; ，j-1 为j为头结点左⼦树节点数量，i-j 为以j为头结点右⼦树节点数量
	//初始化，为了成算，dp[0] = 1,
	//
	//
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i < n+1; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	return dp[n+1]
}

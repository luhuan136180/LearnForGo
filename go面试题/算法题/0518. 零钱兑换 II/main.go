package _518__零钱兑换_II

func change(amount int, coins []int) int {
	//dp
	//初始化
	//确定公式
	//遍历
	dp := make([]int, amount+1)
	dp[0] = 1

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j < amount+1; j++ {
			dp[j] += dp[j-coins[i]]
		}

	}
	return dp[amount]
}
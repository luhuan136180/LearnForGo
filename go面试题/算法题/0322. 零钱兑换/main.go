package main

import (
	"fmt"
	"math"
)

func coinChange(coins []int, amount int) int {
	//3.求最小的状态，及不能初始化为0
	//1.dp[i] :表示amount = i 时的硬币数
	dp := make([]int, amount+1) //0值需要占一个索引位

	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}
	//2.递推公式:dp[j] = min(dp[j-coins[i]+1,dp[j]]
	//4.确定遍历顺序
	for i := 0; i < len(coins); i++ {
		// 遍历背包
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt32 {
				// 推导公式
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
				fmt.Println(dp, j, i)
			}
		}
	}
	//遍历结束
	// 没找到能装满背包的, 就返回-1
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	coins := []int{1, 2, 5}
	a := coinChange(coins, 11)
	fmt.Println(a)
}

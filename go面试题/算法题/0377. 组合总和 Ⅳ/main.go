package _377__组合总和__

//排列问题,不是组合——————————将背包放在外循环，物品放在内循环
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)

	dp[0] = 1

	for i := 0; i < target+1; i++ {
		for j := 0; j < len(nums); j++ {
			if i >= nums[j] {
				dp[i] += dp[i-nums[j]]
			}
		}
	}

	return dp[target]
}

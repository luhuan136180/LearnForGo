package _494__目标和

import "math"

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	if abs(target) > sum {
		return 0
	}

	left := (sum + target) / 2
	if (sum+target)%2 == 1 {
		return 0
	}
	dp := make([]int, left)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := left; j > nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}

	return dp[left]
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

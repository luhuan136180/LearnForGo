package _049__最后一块石头的重量_II

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, v := range stones {
		sum += v
	}
	target := sum / 2

	lenth := len(stones)
	dp := make([]int, target+1)

	for i := 0; i < lenth; i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}

	ans := sum - 2*dp[target]
	return ans
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y

}

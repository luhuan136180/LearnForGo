package _416__分割等和子集

func canPartition(nums []int) bool {
	//1 设置dp[];: i:表示元素的和值
	//2.设置递归公式 dp[i] = max()
	//3.初始化
	//
	sum := 0
	for _, val := range nums {
		sum += val
	}

	dp := make([]int, sum/2+1)

	//dp[0] = 0

	for i := 0; i < len(nums); i++ { //元素（物品）
		for j := len(dp) - 1; j > nums[i]; j-- { //背包重量
			dp[j] = max(dp[j-nums[i]]+nums[i], dp[j])
		}
	}
	return dp[sum/2] == sum/2
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

package _198__打家劫舍

func rob(nums []int) int {

	lenth := len(nums)

	if lenth == 1 {
		return nums[0]
	}
	//定义dp数组
	dp := make([]int, lenth)
	//确定递推公式
	//初始化值
	dp[0] = nums[0]
	dp[1] = max(nums[1], nums[0])
	//确定遍历顺序
	for i := 2; i < lenth; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}

	return dp[lenth-1]
	//推导
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

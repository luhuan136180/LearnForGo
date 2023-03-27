package _1035__不相交的线

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1)+1)
	for i, _ := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}
	ans := 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}

			if ans < dp[i+1][j+1] {
				ans = dp[i+1][j+1]
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

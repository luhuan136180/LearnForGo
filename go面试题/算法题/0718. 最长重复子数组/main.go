package _718__最长重复子数组

func findLength(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1)+1)

	for i, _ := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}

	//
	ans := 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			}
			if ans < dp[i+1][j+1] {
				ans = dp[i+1][j+1]
			}
		}
	}
	return ans
}

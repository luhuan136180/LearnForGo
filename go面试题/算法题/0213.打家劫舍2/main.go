package main

import "fmt"

func rob(nums []int) int {
	//
	lenth := len(nums)
	return max(nums[0]+robnext(nums, 2, lenth-2), robnext(nums, 1, lenth-1))

}

func robnext(nums []int, start, end int) int {
	fmt.Println(start, "_", end)
	if start > end {
		return 0
	}
	dp := make([]int, end-start+1)

	if len(dp) == 1 {
		return nums[start]
	}
	dp[0] = nums[start]
	fmt.Printf("dp[%d]:%d\n", 0, dp[0])
	dp[1] = max(nums[start], nums[start+1])
	fmt.Printf("dp[%d]:%d\n", 1, dp[1])
	for i := 2; i <= len(dp)-1; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i+start])
		fmt.Printf("dp[%d]:%d\n", i, dp[i])
	}
	return dp[len(dp)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{1, 2, 3, 1}
	i := rob(nums)
	fmt.Println("m=", i)
}

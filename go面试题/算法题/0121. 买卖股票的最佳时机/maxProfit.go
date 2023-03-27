package main

import "math"

func maxProfit(prices []int) int {
	mins, maxs := math.MaxInt64, math.MaxInt64
	for i := 0; i < len(prices); i++ {
		//记一个最小值
		mins = min(mins, prices[i])
		//记一个最大值
		maxs = max(maxs, prices[i]-mins)

	}
	//返回最大差值
	return maxs
}
func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

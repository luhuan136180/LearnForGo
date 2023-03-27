package main

import "math"

func maxProfit(prices []int) int {
	ans := 0
	for i := 1; i < len(prices); i++ {
		ans += int(math.Max(0, float64(prices[i]-prices[i-1])))
	}
	return ans
}

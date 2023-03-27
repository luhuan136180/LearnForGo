package main

import (
	"math"
	"sort"
)

func largestSumAfterKNegations(nums []int, K int) int {
	//Slice对切片x进行排序
	sort.Slice(nums, func(i, j int) bool {
		//math.Abs:返回绝对值
		// // Less方法报告索引i的元素是否比索引j的元素小
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})

	//将复数转为正数
	for i := 0; i < len(nums); i++ {
		if K > 0 && nums[i] < 0 {
			nums[i] = -nums[i]
			K--
		}
	}
	//从最小的正数改为负数
	if K%2 == 1 {
		nums[len(nums)-1] = -nums[len(nums)-1]
	}

	result := 0
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

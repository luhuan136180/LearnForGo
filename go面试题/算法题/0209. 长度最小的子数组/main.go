package _209__长度最小的子数组

import "math"

func minSubArrayLen(target int, nums []int) int {
	//滑动数组，以尾指针作为遍历基础
	//动态调整首指针，计算滑动窗口长度
	lenth := len(nums)

	result := math.MaxInt64
	left, sum := 0, 0
	for right := 0; right < lenth; right++ {
		sum += nums[right]
		for sum >= target {
			tlen := right - left + 1
			if result > tlen {
				result = tlen
			}
			sum -= nums[left]
			left++

		}
	}

	if result == math.MaxInt64 {
		return 0
	}

	return result
}

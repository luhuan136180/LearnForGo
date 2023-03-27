package main

func maxSubArray(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		count := 0
		for j := i; j < len(nums); j++ {
			count += nums[j]
			if result < count {
				result = count
			}
		}

	}
	return result
}

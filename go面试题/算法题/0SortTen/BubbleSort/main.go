package main

func bubbleSort(nums []int) []int {
	n := len(nums)
	for i := 0; i < len(nums)-1; i++ {
		exchange := false
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				exchange = true
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
		if !exchange {
			return nums
		}
	}

}

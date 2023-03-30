package main

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	temp := len(nums)
	for left <= right {
		mid := ((right - left) >> 1) + left //等同于：(right-left)%2
		if target <= nums[mid] {
			temp = mid
			right = mid - 1
		} else {

			left = mid + 1
		}

	}
	return temp
}

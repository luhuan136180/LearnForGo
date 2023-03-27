package _704__二分查找

func search(nums []int, target int) int {
	//首先设定左右节点
	left := 0
	right := len(nums) - 1

	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] > target {
			right = middle - 1
		} else if nums[middle] < target {
			left = middle + 1
		} else if nums[middle] == target {
			return middle
		}
	}
	return -1
}

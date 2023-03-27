package _034__在排序数组中查找元素的第一个和最后一个位置

func searchRange(nums []int, target int) []int {
	leftBorder := getLeftBorder(nums, target)
	rightBorder := getRightBorder(nums, target)

	if leftBorder == -2 || rightBorder == -2 {
		return []int{-1, -1}
	}
	if rightBorder-leftBorder > 1 {
		return []int{leftBorder + 1, rightBorder - 1}
	}
	return []int{-1, -1}
}

func getRightBorder(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	rightBorder := -2 // 记录一下rightBorder没有被赋值的情况

	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
			rightBorder = left
		}
	}
	return rightBorder
}

func getLeftBorder(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	leftBordor := -2
	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] >= target {
			right = middle - 1
			leftBordor = right
		} else {
			left = middle + 1
		}
	}
	return leftBordor
}

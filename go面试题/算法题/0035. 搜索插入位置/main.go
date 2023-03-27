package _035__搜索插入位置

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] < target {
			left = middle + 1
		} else if nums[middle] > target {
			right = middle - 1
		} else if nums[middle] == target {
			return middle
		}

	}

	// 分别处理如下四种情况
	// 目标值在数组所有元素之前 [0,0)
	// 目标值等于数组中某一个元素 return middle
	// 目标值插入数组中的位置 [left, right) ，return right 即可
	// 目标值在数组所有元素之后的情况 [left, right)，因为是右开区间，所以 return right
	return right + 1
}

package _977_有序数组的平方

import "sort"

//暴力
func sortedSquares(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		temp := nums[i]
		nums[i] = temp * temp
	}
	sort.Ints(nums)
	return nums
}

//双指针法，从两头向中间遍历
func sortedSquares2(nums []int) []int {

	left := 0
	right := len(nums) - 1

	str := right
	ans := make([]int, len(nums))
	for left <= right {
		if nums[left]*nums[left] > nums[right]*nums[right] {
			ans[str] = nums[left] * nums[left]
			str--
		} else {
			ans[str] = nums[right] * nums[right]
			str--
		}
	}

	return ans
}
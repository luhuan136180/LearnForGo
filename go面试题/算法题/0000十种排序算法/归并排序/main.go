package main

import "fmt"

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	middle := len(nums) / 2
	left := mergeSort(nums[:middle])
	right := mergeSort(nums[middle:])

	result := merge(left, right)
	return result
}
func merge(left, right []int) []int {
	res := make([]int, 0)
	l, r := 0, 0
	llen, rlen := len(left), len(right)

	for l < llen || r < rlen {
		if l == llen {
			res = append(res, right[r:]...)
			break
		}
		if r == rlen {
			res = append(res, left[l:]...)
			break
		}
		if left[l] < right[r] {
			res = append(res, left[l])
			l++
		} else {
			res = append(res, right[r])
			r++
		}
	}
	return res
}

func main() {
	arr := []int{8, 9, 5, 7, 3, 5, 2, 6, 1, 4}
	res := mergeSort(arr)
	fmt.Println("arr=", res)
}

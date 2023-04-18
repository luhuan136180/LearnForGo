package 选择排序

import "fmt"

//比较拗口，举个例子，序列5 8 5 2 9， 我们知道第一遍选择第1个元素5会和2交换，那么原序列中2个5的相对前后顺序就被破坏了，所以选择排序不是一个稳定的排序算法。

func SelectSort(nums []int, n int) {
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			tmp := nums[minIndex]
			nums[minIndex] = nums[i]
			nums[i] = tmp
		}
		fmt.Println(nums)
	}
	fmt.Println(nums)
}

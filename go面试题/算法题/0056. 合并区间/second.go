package main

import "sort"

func merge2(intervals [][]int) [][]int {
	res := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	left := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			intervals[i][1] = max(intervals[i][1], intervals[i-1][1])

		} else {
			ans := []int{left, intervals[i-1][1]}
			res = append(res, ans)
			left = intervals[i][0]
		}
	}

	return res
}

//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}

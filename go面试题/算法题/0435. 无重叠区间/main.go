package _435__无重叠区间

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	count := 0
	sort.Slice(intervals, func(i, j int) bool {

		return intervals[i][0] < intervals[j][0]
	})

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			count++
			intervals[i][1] = min(intervals[i][1], intervals[i-1][1])
		}
	}
	return count
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

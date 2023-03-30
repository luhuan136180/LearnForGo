package main

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]

	})
	res := 1
	for i := 1; i < len(points); i++ {
		if points[i][0] < points[i-1][1] {
			points[i][1] = min(points[i][1], points[i-1][1])
		} else {
			res++
		}
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

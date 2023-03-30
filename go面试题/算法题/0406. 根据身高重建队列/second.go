package main

import "sort"

func reconstructQueue2(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	res := make([][]int, 0)
	//排序完成
	for i := 0; i < len(people); i++ {
		index := people[i][1]
		res = append(res, people[i])
		copy(res[index+1:], res[index:])
		res[index] = people[i]
	}
	return res
}

package _090__子集_II

import (
	"sort"
)

var (
	res  [][]int
	path []int
	used []bool //下标对应nums的下标
)

func subsetsWithDup(nums []int) [][]int {
	res, path = make([][]int, 0), make([]int, 0)
	used = make([]bool, len(nums))
	sort.Ints(nums)
	dfs(nums, 0)
	return res
}

func dfs(nums []int, index int) {
	tmp := make([]int, len(path))
	copy(tmp, path)
	res = append(res, tmp)

	for i := index; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && used[i-1] == false {
			continue
		}
		path = append(path, nums[i])
		used[i] = true
		dfs(nums, i+1)
		path = path[:len(path)-1]
		used[i] = false
	}
}

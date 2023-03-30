package _047__全排列_II

import "sort"

var (
	res  [][]int
	path []int
	used []bool
)

func permuteUnique(nums []int) [][]int {
	res, path, used = make([][]int, 0), make([]int, 0), make([]bool, len(nums))
	sort.Ints(nums)
	dfs(nums, 0)
	return res
}
func dfs(nums []int, length int) {
	if length == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
	}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && used[i-1] == false {
			//数层去重
			continue
		}
		if !used[i] {
			path = append(path, nums[i])
			used[i] = true
			dfs(nums, length+1)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
}

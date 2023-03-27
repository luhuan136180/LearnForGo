package _040__组合总和_II

import "sort"

var (
	res  [][]int
	path []int
	used []bool
)

func combinationSum2(candidates []int, target int) [][]int {
	res, path = make([][]int, 0), make([]int, 0)
	used = make([]bool, len(candidates))
	sum := 0
	start := 0
	sort.Ints(candidates) // 排序，为剪枝做准备
	dfs(candidates, target, sum, start)
	return res
}

func dfs(candidates []int, target, sum, start int) {
	if sum == target {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}

	for i := start; i < len(candidates); i++ {
		if sum+candidates[i] > target {
			break
		}
		// used[i - 1] == true，说明同一树枝candidates[i - 1]使用过
		// used[i - 1] == false，说明同一树层candidates[i - 1]使用过
		if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == false {
			continue
		}

		sum += candidates[i]
		used[i] = true
		path = append(path, candidates[i])
		dfs(candidates, target, sum, i+1)
		used[i] = false
		path = path[:len(path)-1]
		sum -= candidates[i]
	}
}

package _046__全排列

var (
	res  [][]int
	path []int
	used []bool
)

func permute(nums []int) [][]int {
	res, path, used = make([][]int, 0), make([]int, 0), make([]bool, len(nums))
	dfs(nums, 0)
	return res
}

func dfs(nums []int, lenth int) {
	if len(nums) == lenth {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
	}

	for i := 0; i < len(nums); i++ {
		if used[i] == true {
			continue
		}
		path = append(path, nums[i])
		lenth++
		used[i] = true
		dfs(nums, lenth)
		lenth--
		path = path[:len(path)-1]
		used[i] = false
	}
}

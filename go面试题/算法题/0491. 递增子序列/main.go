package _491__递增子序列

var (
	res  [][]int
	path []int
)

func findSubsequences(nums []int) [][]int {
	res, path = make([][]int, 0), make([]int, 0)
	index := 0
	length := 0
	dfs(nums, index, length)
	return res
}

//不能对数组排序，去重逻辑要修改
func dfs(nums []int, index, length int) {
	if length >= 2 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
	}
	//将显示重复的逻辑放在每一层单独做
	//去重的核心思想是同层去重
	used := make(map[int]bool, len(nums))
	for i := index; i < len(nums); i++ {
		if used[nums[i]] { //去重
			continue //该分支会和之前的结果集重复，退出
		}

		if len(path) == 0 || nums[i] >= path[len(path)-1] {
			path = append(path, nums[i])
			length++
			used[nums[i]] = true //进入循环前将user【i】置位已使用
			dfs(nums, i+1, length)

			length--
			path = path[:len(path)-1]
		}
	}
}

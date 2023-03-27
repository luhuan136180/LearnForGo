package _216__组合总和_III

var (
	res  [][]int
	path []int
)

func combinationSum3(k int, n int) [][]int {
	res = make([][]int, 0)
	path = make([]int, 0)
	sum := 0
	start := 1
	dfs(n, k, start, sum)
	return res
}

func dfs(n, k, start, sum int) {
	if len(path) == k {
		//已经递归（循环）k次，当前为k个数相加
		if sum == n {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)

		}

		return
	}

	for i := start; i <= 9; i++ {
		sum += i
		if sum > n || 9-i+1 < k-len(path) { //剪枝
			break
		}
		path = append(path, i)
		dfs(n, k, i+1, sum) //递归

		//回溯，一共需要两步
		sum -= i
		path = path[:len(path)-1]
	}
}

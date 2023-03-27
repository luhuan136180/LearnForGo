package _039__组合总和

//如果是一个集合来求组合的话，就需要startIndex，
//多个集合取组合，各个集合之间相互不影响，那么就不用startIndex，
var (
	res  [][]int
	path []int
)

func combinationSum(candidates []int, target int) [][]int {
	res, path = make([][]int, 0), make([]int, 0)
	start := 0
	sum := 0
	dfs(candidates, target, sum, start)
	return res
}

func dfs(candidates []int, target, sum, start int) {
	if sum >= target {
		if sum == target { //找到一个结果
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)

		}

		return
	}

	for i := start; i < len(candidates); i++ { //从当前索引的元素是可以使用的
		sum += candidates[i]
		path = append(path, candidates[i])

		dfs(candidates, target, sum, i)
		path = path[:len(path)-1] //回溯
		sum -= candidates[i]      //回溯
	}

}

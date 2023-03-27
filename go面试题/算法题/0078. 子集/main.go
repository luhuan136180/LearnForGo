package _078__子集

var (
	res  [][]int
	path []int
)

func subsets(nums []int) [][]int {
	res, path = make([][]int, 0), make([]int, 0)

	dfs(nums, 0)
	return res
}

func dfs(nums []int, index int) {
	tmp := make([]int, len(path))
	copy(tmp, path)
	res = append(res, tmp)
	//if index == len(nums) {
	//	return
	//}

	for i := index; i < len(nums); i++ {
		path = append(path, nums[i])
		dfs(nums, i+1)            //递归
		path = path[:len(path)-1] //回溯
	}
}

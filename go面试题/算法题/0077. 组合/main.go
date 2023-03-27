package _077__组合

var (
	path []int
	res  [][]int
)

//k:递归的层数
func combine(n int, k int) [][]int {
	//初始化
	res = make([][]int, 0)
	path = make([]int, 0)
	dfs(n, k, 1)
	return res
}

func dfs(n, k, start int) {
	if len(path) == k {
		//的到想要的一种答案
		temp := make([]int, k)
		copy(temp, path)
		res = append(res, temp)
		return //回退
	}

	for i := start; i <= n; i++ {
		//当剩余集合已经不足以用于递归的消耗时，不可能有需要打的结果
		if n-i+1 < k-len(path) { // 剪枝
			break
		}
		path = append(path, i)
		dfs(n, k, i+1)            //递归
		path = path[:len(path)-1] //回溯
	}
}

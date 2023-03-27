package _300__最长递增子序列

func lengthOfLIS(nums []int) (ans int) {
	n := len(nums)
	memo := make([]int, n) // 本题可以初始化成 0，表示没有计算过
	var dfs func(int) int
	dfs = func(i int) int {
		p := &memo[i] //指针变量
		if *p > 0 {
			return *p
		}
		//*p = 0
		res := 0
		for j, x := range nums[:i] {
			//j:索引 ;;; x:nums的值
			if x < nums[i] { //
				res = max(res, dfs(j)) //找最长的子序列
			}
		}
		res++ //加上自己后的子序列长度
		*p = res
		return res
	}

	for i := 0; i < n; i++ { //从左边开始寻找最长
		ans = max(ans, dfs(i))
	}
	return

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

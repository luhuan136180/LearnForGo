package _131__分割回文串

var (
	path []string
	res  [][]string
)

func partition(s string) [][]string {
	res, path = make([][]string, 0), make([]string, 0)
	start := 0
	dfs(s, start)
	return res
}

func dfs(s string, startindex int) {
	//切割线切到了字符串最后面，说明找到了一种切割方法
	if startindex == len(s) {
		tmp := make([]string, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}

	for i := startindex; i < len(s); i++ {
		str := s[startindex : i+1]
		if isPalindrome(str) == true { //确定为回文字
			path = append(path, str)
			dfs(s, i+1)
			path = path[:len(path)-1]
		}
		//不是回文字串的话就继续遍历
	}
}
func isPalindrome(s string) bool {
	start, end := 0, len(s)-1
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

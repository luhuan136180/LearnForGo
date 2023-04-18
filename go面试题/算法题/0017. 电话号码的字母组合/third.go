package _017__电话号码的字母组合

import "strconv"

var (
	res  []string
	path []byte
	dir  []string
)

func letterCombinations(digits string) []string {
	dir = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	path, res = make([]byte, 0), make([]string, 0)
	dfs(digits, 0)
	return res
}

func dfs(digits string, index int) {
	if len(path) == len(digits) {

		ans := string(path)
		res = append(res, ans)
		return
	}
	num, _ := strconv.Atoi(string(digits[index]))
	str := dir[num]
	for i := 0; i < len(str); i++ {
		path = append(path, str[i])
		dfs(digits, index+1)
		path = path[:len(path)-1]
	}
}

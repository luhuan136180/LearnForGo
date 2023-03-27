package _022__括号生成

var (
	res  []string
	path []byte
)

func generateParenthesis(n int) []string {
	m := n * 2
	res, path = make([]string, 0), make([]byte, 0)
	i, open := 0, 0
	dfs(i, open, m)
	return res
}
func dfs(i, open int, m int) {
	if i == m {
		tmp := make([]byte, len(path))
		copy(tmp, path)
		res = append(res, string(tmp))
		return
	}
	if open < m/2 {
		path = append(path, '(')
		dfs(i+1, open+1, m)
		path = path[:len(path)-1]
	}
	if i-open < open {
		path = append(path, ')')
		dfs(i+1, open, m)
		path = path[:len(path)-1]
	}
}

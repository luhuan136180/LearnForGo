package _017__电话号码的字母组合

var (
	res2  []string
	a     []string
	path2 []byte
)

func letterCombinations2(digits string) []string {
	a = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	dfs(digits, 0)
	return res2
}

func dfs2(digits string, index int) {
	if index == len(digits) {
		ans := string(path2)
		res = append(res2, ans)
	}
	number := int(digits[index] - '0')
	str := a[number-2]
	for _, val := range str {
		path2 = append(path2, byte(val))
		dfs(digits, index+1)
		path2 = path2[:len(path2)-1]
	}

}

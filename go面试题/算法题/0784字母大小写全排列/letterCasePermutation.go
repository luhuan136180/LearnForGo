package main

import (
	"fmt"
	"unicode"
)

func main() {
	ans := letterCasePermutation2("Abcd")
	fmt.Println(ans)
}

//方法一：广度优先搜索
func letterCasePermutation(s string) []string {
	var ans []string
	q := []string{""}
	for len(q) > 0 {
		cur := q[0]
		pos := len(cur)
		if pos == len(s) {
			ans = append(ans, cur)
			q = q[1:]
		} else {
			if unicode.IsLetter(rune(s[pos])) {
				q = append(q, cur+string(s[pos]^32))
			}
			q[0] += string(s[pos])
		}

	}
	return ans

}

//方法二：回溯
func letterCasePermutation2(s string) (ans []string) {

	t := []byte(s)
	var dfs func(int)
	dfs = func(i int) {
		if i >= len(t) {
			ans = append(ans, string(t))
			return
			
		}
		dfs(i + 1)
		if t[i] >= 'A' {
			t[i] ^= 32
			dfs(i + 1)
		}
	}

	dfs(0)
	return ans
}

//普通版的二插遍历
func letterCasePermutation3(s string) []string {
	n := len(s)
	str := []rune(s)
	var ans []string
	var dfs func(cur int)
	dfs = func(cur int) {
		if cur == n {
			ans = append(ans, string(str))
			return
		}
		if unicode.IsLetter(str[cur]) {
			str[cur] = unicode.ToLower(str[cur])
			dfs(cur + 1)
			str[cur] = unicode.ToUpper(str[cur])
			dfs(cur + 1)
		} else { //数字
			dfs(cur + 1)
		}
	}
	dfs(0)
	return ans
}

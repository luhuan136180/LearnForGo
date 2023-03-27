package _132__分割回文串_II

import "math"

var depth int

func minCut(s string) int {
	depth = math.MaxInt64
	dfs(s, 0, 0)
	return depth
}

func dfs(s string, startindex, dep int) {
	if startindex == len(s) {
		if dep < depth {
			depth = dep
		}
	}
	if dep >= depth {
		return
	}
	for i := startindex; i < len(s); i++ {

		str := s[startindex : i+1]
		if isPalindrome(str) {
			dfs(s, i+1, dep+1)
		}
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

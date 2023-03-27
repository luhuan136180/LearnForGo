package main

func partitionLabels(s string) []int {
	var res []int
	var marks [26]int
	size, left, right := len(s), 0, 0
	for i := 0; i < size; i++ {
		//遍历统计出s字符串中出现的各个字母在字符串的最远距离
		marks[s[i]-'a'] = i
	}
	for i := 0; i < size; i++ {
		right = max(right, marks[s[i]-'a'])
		if i == right {
			res = append(res, right-left+1)
			left = i + 1
		}
	}
	return res
}

func max(a int, b int) int {
	if a < b {
		a = b
	}
	return a
}

func partitionLabels2(s string) []int {
	var res []int
	var mark [26]int
	size, left, right := len(s), 0, 0
	for i := 0; i < size; i++ {
		mark[s[i]-'a'] = i
	}

	for i := 0; i < size; i++ {
		right = max(right, mark[s[i]-'a'])
		if i == right {
			res = append(res, right-left+1)
			left = i + 1
		}
	}
	return res
}

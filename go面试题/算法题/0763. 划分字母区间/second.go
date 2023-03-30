package main

func partitionLabels3(s string) []int {
	res := make([]int, 0)
	mark := make(map[byte]int, 0)
	for i, _ := range s {
		if i > mark[s[i]] {
			mark[s[i]] = i
		}
	}

	left, right := 0, mark[s[0]]
	for i := 0; i < len(s); i++ {
		curindex := mark[s[i]]
		if curindex > right {
			right = curindex
		}
		if i == right {
			res = append(res, right-left+1)
			left = i + 1
		}

	}
	return res
}

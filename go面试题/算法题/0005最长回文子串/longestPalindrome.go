package main

func longestPalindrome(s string) string {
	if s == "" {
		//字符串为空
		return ""
	}
	start, end := 0, 0 //记录回文字子串的起点和终点
	maxlen := 1
	for i := 0; i < len(s); i++ {
		lenth := 0
		left := i - 1  //左节点
		right := i + 1 //右节点
		//向左侧拓展,直到超过边界或遇到与中心字符不等跳出while循环
		for left >= 0 && s[left] == s[i] {
			left--
			lenth++
		}
		//向右侧扩展，直到超过边界或遇到与中心字符不等跳出while循环
		for right <= len(s)-1 && s[right] == s[i] {
			right++
			lenth++
		}
		//同时向两边拓展
		for left >= 0 && right <= len(s)-1 && s[left] == s[right] {
			left--
			right++
			lenth += 2
		}
		if lenth > maxlen {
			maxlen = lenth
			start = left
			end = right

		}

	}
	return s[start+1 : end]
}

func longestPalindrome2(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	maxlen := 1
	len := len(s)

	for i := 0; i < len; i++ {
		left := i - 1
		right := i + 1
		lenth := 0
		if left >= 0 && s[left] == s[i] {
			left--
			lenth++
		}
		if right < len && s[right] == s[i] {
			right++
			lenth++
		}
		if right < len && left >= 0 && s[right] == s[left] {
			left--
			right++
			lenth++
		}

		if lenth > maxlen {
			maxlen = lenth
			start = left
			end = right
		}

	}

	return s[start+1 : end]

}

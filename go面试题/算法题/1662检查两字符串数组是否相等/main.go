package main

import "strings"

//暴力解法
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var s1, s2 string
	for _, val := range word1 {
		s1 += val
	}
	for _, val := range word2 {
		s2 += val
	}
	if s1 == s2 {
		return true
	}
	return false
}

//简化版暴力
func arrayStringsAreEqual2(word1 []string, word2 []string) bool {
	//func Join(elems []string, sep string) string
	return strings.Join(word1, "") == strings.Join(word2, "")

}

//遍历
func arrayStringsAreEqual3(word1 []string, word2 []string) bool {
	//设置四个指针，p1，p2用于指向两数组的元素，ij指向string元素的每个字符
	var i, j, p1, p2 int
	for p1 < len(word1) && p2 < len(word2) {
		//遍历完两数组，
		//先比较起始位
		if word1[p1][i] != word2[p2][j] { //比较两数组第一个元素的第一个字符
			return false
		}
		i++
		if i == len(word1[p1]) { //i指针遍历p1指向元素的字符
			//该元素遍历结束
			i = 0
			p1++ //p1指向下个元素
		}
		j++
		if j == len(word2[p2]) {
			j = 0
			p2++
		}

	}
	//遍历结束，及遍历过程中没有发现不相同得到字符
	return p1 == len(word1) && p2 == len(word2)

}

func arrayStringsAreEqual4(word1 []string, word2 []string) bool {
	var i, j, p1, p2 int
	for p1 < len(word1) && p2 < len(word2) {
		if word1[p1][i] != word2[p2][j] {
			return false
		}
		i++
		if i == len(word1[p1]) {
			i = 0
			p1++
		}

		j++
		if j == len(word2[p2]) {
			j = 0
			p2++
		}

	}
	return p1 == len(word1) && p2 == len(word2)
}

package main

import "sort"

/*假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。
对每个孩子 i，都有一个胃口值 g[i]，这是能让孩子们满足胃口的饼干的最小尺寸；
并且每块饼干 j，都有一个尺寸 s[j] 。如果 s[j] >= g[i]，
我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足。
你的目标是尽可能满足越多数量的孩子，并输出这个最大数值。*/
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	num := 0

	for _, val := range s {
		if num < len(g) && g[num] <= val {
			num++
		}
	}
	return num
}

func findContentChildren2(g []int, s []int) (ans int) {
	sort.Ints(g)
	sort.Ints(s)
	lenS := len(s)
	j := 0
	for _, val := range g {
		for j < lenS && val > s[j] {
			j++
		}
		if j < lenS { //找到了
			ans++
			j++
		}
	}
	return
}

func findContentChildren3(g []int, s []int) (ans int) {
	sort.Ints(g) //小孩列表
	sort.Ints(s)
	lens := len(s)
	j := 0
	for _, val := range g {
		for j < lens && val > s[j] {
			j++
		}
		if j < lens {
			ans++
			j++
		}
	}

	return
}

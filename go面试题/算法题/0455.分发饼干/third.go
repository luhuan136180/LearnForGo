package main

import "sort"

func findContentChildren4(g []int, s []int) (ans int) {
	sort.Ints(g)
	sort.Ints(s)
	child := 0
	for i := 0; i < len(s); i++ { //小饼干先喂食量小的
		if child < len(g) && g[child] <= s[i] {
			//当前饼干可以满足这个小孩
			//	吃饱小孩+1；饼干向后遍历
			child++
		}

	}
	ans = child
	return
}

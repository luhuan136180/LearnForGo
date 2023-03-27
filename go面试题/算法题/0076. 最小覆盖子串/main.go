package _076__最小覆盖子串

import "math"

func minWindow(s string, t string) string {
	//滑动窗口,
	//用map存储t中各个字符出现的次数
	//因为需要返回最小字符串，所以初始化一个	R,L,lenth记录最小字符串的右，左，长度

	Lprt, Rprt, lenth := 0, 0, math.MaxInt64
	//两个map，一个是t的Tmap，用于统计数显的字符数，和该字符个数
	//第二个map  Smap用于统计滑窗中的字符和其个数，当Smap中完全包含 tmap，存在
	Tmap, Smap := make(map[byte]int), make(map[byte]int)

	for i := 0; i < len(t); i++ {
		Tmap[t[i]]++
	}

	left := 0
	for i := 0; i < len(s); i++ {
		Smap[s[i]]++
		for check(Tmap, Smap) {
			curlen := i - left + 1
			if lenth > curlen {
				lenth = curlen
				Rprt = i
				Lprt = left
			}
			Smap[s[left]]--

			left++
		}
	}

	//lenth = math.MaxInt64 表示没有找到
	if lenth == math.MaxInt64 {
		return ""
	}
	return s[Lprt : Rprt+1]
}

func check(tmap, smap map[byte]int) bool {
	for i, val := range tmap {
		if smap[i] < val {
			return false
		}
	}
	return true
}

package main

import "math"

func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{} //key：左边界起始字符，value：不重复的字符串
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ { //外循环用于移动滑窗的左边界

		if i != 0 {
			//左指针向右移动一格，移除一个字符
			delete(m, s[i-1]) //从m哈希列表中移出key=s[i-1]的键值对
		}

		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		//
		ans = int(math.Max(float64(ans), float64(rk-i+1)))
	}
	return ans
}

func lengthOfLongestSubstring2(s string) int {
	var n = len(s)
	if n <= 1 {
		return n
	}
	MaxLen := 1
	var left, right, window = 0, 0, make(map[byte]bool)
	for right < n {
		var rightChar = s[right]
		if window[rightChar] { //若window中存在rightChar字符，则返回true
			//滑窗向右移动边界
			delete(window, s[left])
			left++
		}

		if right-left+1 > MaxLen { //每次移动有边界后，比较滑窗大小

			MaxLen = right - left + 1
		}
		window[rightChar] = true //将新的右边界的字符设为true
		right++
	}
	return MaxLen
}

func lengthOfLongestSubstring3(s string) int {
	var n = len(s)
	if n < 1 {
		return n
	}
	var MaxLen = 1
	var left, right, window = 0, 0, make(map[byte]int)
	for right < n {
		var rightChar = s[right]
		var rightCharIndex = 0
		if _, ok := window[rightChar]; ok {
			rightCharIndex = window[rightChar]
		}
		if rightCharIndex > left {
			left = rightCharIndex
		}
		if right-left+1 > MaxLen {
			MaxLen = right - left + 1
		}
		window[rightChar] = right + 1
		right++
	}
	return MaxLen
}

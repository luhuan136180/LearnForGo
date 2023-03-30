package main

import "strconv"

func monotoneIncreasingDigits2(n int) int {
	//为了方便遍历，将int型转为string型
	str := strconv.Itoa(n)
	ss := []byte(str)
	if len(ss) <= 1 {
		return n
	}
	for i := len(ss) - 2; i >= 0; i-- {
		if ss[i] > ss[i+1] {
			ss[i] -= 1
			for j := i + 1; j < len(ss); j++ {
				ss[j] = '9'
			}
		}
	}
	res, _ := strconv.Atoi(string(ss))
	return res
}

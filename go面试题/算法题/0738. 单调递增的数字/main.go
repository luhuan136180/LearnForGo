package main

import "strconv"

func monotoneIncreasingDigits(n int) int {
	s := strconv.Itoa(n) //将数字转换为字符串，方便使用下标
	ss := []byte(s)      //将字符串转换为byte数组，方便更改
	size := len(ss)

	if size <= 1 {
		return n
	}

	for i := n - 1; i > 0; i-- {
		//前一个大于后一位,前一位减1，后面的全部置为9
		if ss[i-1] > ss[i] {
			ss[i-1] -= 1
			for j := i; j < n; j++ {
				ss[j] = '9'
			}
		}
	}
	res, _ := strconv.Atoi(string(ss))
	return res
}

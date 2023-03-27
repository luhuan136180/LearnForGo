package main

import "math"

func myAtoi(s string) int {
	sign, result, n, i := 1, 0, len(s), 0
	//const MinInt32, MaxInt32 = -1 << 31, 1<<31 - 1
	for index, val := range s {
		if val != ' ' {
			i = index
			break
		}

	}
	if i >= n {
		return 0
	}
	switch s[i] {
	case '-':
		sign = -1
		i++
	case '+':
		i++
	}

	for ; i < n && s[i] >= '0' && s[i] <= '9'; i++ {
		result = result*10 + int(s[i]-'0')
		if sign*result > math.MaxInt32 {
			return math.MaxInt32
		} else if sign*result < math.MinInt32 {
			return math.MinInt32
		}
	}
	return result * sign
}

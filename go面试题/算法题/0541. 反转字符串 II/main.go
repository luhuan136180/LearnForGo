package _541__反转字符串_II

func reverseStr(s string, k int) string {
	ss := []byte(s)
	lenth := len(s)

	for i := 0; i < lenth; i += k * 2 {
		if i+k <= lenth {
			reverse(ss[i : i+k])
		} else {
			reverse(ss[i:])
		}
	}
	return string(ss)
}

func reverse(b []byte) {
	left := 0
	right := len(b) - 1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

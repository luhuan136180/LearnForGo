package main

func arraySign(nums []int) int {
	n := 0
	for _, val := range nums {
		if val == 0 {
			return 0
		} else if val < 0 {
			n++
		}
	}
	if n%2 == 1 {
		return -1
	} else {
		return 1
	}
}

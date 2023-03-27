package main

func reverse(x int) int {
	res := 0
	for x != 0 {
		y := x % 10
		res = res*10 + y
		if !(res <= (1<<31)-1 && -(1<<31) <= res) {
			return 0
		}
		x = x / 10
	}
	return res
}

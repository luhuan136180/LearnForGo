package main

func candy(ratings []int) int {
	candys := make([]int, len(ratings))
	for i := 1; i < len(ratings); i++ {
		if ratings[i-1] < ratings[i] {
			candys[i] = candys[i-1] + 1
		}
	}
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candys[i] = max(candys[i], candys[i+1]+1)
		}
	}
	res := 0
	for _, val := range candys {
		res += val
	}
	return res + len(ratings)
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

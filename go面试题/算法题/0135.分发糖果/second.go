package main

func candy2(ratings []int) int {
	ans := 0
	res := make([]int, len(ratings))
	res[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			res[i] = res[i-1] + 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			res[i] = max(res[i], res[i+1]+1)
		}
	}

	for _, val := range res {
		ans += val
	}
	return ans
}

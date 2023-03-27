package main

import "fmt"

//// 双指针解法：
func trap(height []int) int {
	//确定两个表
	lmaxHeight := make([]int, len(height))
	rmaxHright := make([]int, len(height))

	//初始化
	for i, val := range height {
		if i == 0 {
			lmaxHeight[0] = val
		} else {
			lmaxHeight[i] = max(lmaxHeight[i-1], val)
		}
	}

	for j := len(height) - 1; j >= 0; j-- {
		if j == len(height)-1 {
			rmaxHright[j] = height[j]
		} else {
			rmaxHright[j] = max(height[j], height[j+1])
		}
	}

	fmt.Println("lm", lmaxHeight)
	fmt.Println("rm", rmaxHright)

	sum := 0

	for i := 0; i < len(height); i++ {
		count := min(lmaxHeight[i], rmaxHright[i]) - height[i]
		if count > 0 {
			sum += count
		}
	}
	return sum

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func main() {
	a := []int{4, 2, 0, 3, 2, 5}
	trap(a)
}
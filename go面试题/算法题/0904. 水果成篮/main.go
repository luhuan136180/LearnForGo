package _904__水果成篮

func totalFruit(fruits []int) int {
	//使用map存储篮子，k果子类型，val数量
	//使用滑动窗口解题,此时滑动窗口大小表示果子树量
	//滑动是双循环法，内循环用于找首指针的前进距离
	a := make(map[int]int)
	ans := 0
	left := 0
	for i := 0; i < len(fruits); i++ {
		a[fruits[i]]++
		for len(a) > 2 {
			a[fruits[left]]--
			if a[fruits[left]] == 0 {
				delete(a, fruits[left])
			}
			left++
		} //遍历结束代表a中只有两个或一个果子种类了
		ans = max(ans, i-left+1)
	}

	return ans
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

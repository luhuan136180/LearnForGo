package _739__每日温度

//暴力解法——超时了
func dailyTemperatures(temperatures []int) []int {
	lenth := len(temperatures)
	ans := make([]int, lenth)
	for i := 0; i < lenth-1; i++ {
		for j := i + 1; j < lenth; j++ {
			if temperatures[i] < temperatures[j] {
				ans[i] = j - i
				break
			}

		}
	}
	ans[lenth-1] = 0
	return ans
}

//单调栈函数
func dailyTemperatures2(temperatures []int) []int {
	lenth := len(temperatures)
	ans := make([]int, lenth)

	//制作一个栈
	stack := []int{}
	for i, val := range temperatures {
		//当栈里面有元素，而且即将进入的元素大于栈顶元素是，弹出栈顶元素
		//计算下标差值
		for len(stack) >= 0 && val > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			ans[i] = i - top
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	return ans
}

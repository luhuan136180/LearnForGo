package _739__每日温度

func dailyTemperatures3(num []int) []int {
	res := make([]int, len(num))
	stack := make([]int, 0)

	for i := 0; i < len(num); i++ {
		//双层for
		for len(stack) > 0 && num[i] > num[stack[len(stack)-1]] {
			res[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}

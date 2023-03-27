package _167__两数之和_II___输入有序数组

func twoSum(numbers []int, target int) []int {
	//用key存元素大小，val存下标用于找位置
	ans := []int{}
	set := make(map[int]int, 0)
	for i, v := range numbers {

		if _, ok := set[target-v]; ok {
			//ok = true ,则存在两个元素
			ans = append(ans, set[target-v]+1)
			ans = append(ans, i+1)
		} else {
			set[v] = i
		}
	}
	return ans
}

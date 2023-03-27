package main

func twoSum2(nums []int, target int) []int {
	//用key存元素大小，val存下标用于找位置
	ans := []int{}
	set := make(map[int]int, 0)
	for i, v := range nums {

		if _, ok := set[target-v]; ok {
			//ok = true ,则存在两个元素
			ans = append(ans, set[target-v])
			ans = append(ans, i)
		} else {
			set[v] = i
		}
	}
	return ans
}

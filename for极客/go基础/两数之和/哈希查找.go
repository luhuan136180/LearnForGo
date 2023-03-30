package main

// 时间复杂度：O(n)
// 空间复杂度：O(n)
// 空间换时间

func main() {

}
func twoSum2(nums []int, target int) []int {
	//数据预处理
	val := make(map[int]int)
	for i, _ := range nums {
		val[nums[i]] = i
	}

	for i, _ := range nums {
		x := nums[i]
		if index, ok := val[target-x]; ok {
			if index != i {
				return []int{i, index}
			}
		}
	}
	return []int{}
}

func twoSum(nums []int, target int) []int {
	var numIndexMap = make(map[int]int)
	for i := range nums {
		var x = nums[i]
		// 2. 哈希查找 - O(1)
		if index, ok := numIndexMap[target-x]; ok {
			return []int{i, index}
		}
		numIndexMap[x] = i
	}
	return []int{}
}

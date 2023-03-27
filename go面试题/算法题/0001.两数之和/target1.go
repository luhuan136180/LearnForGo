package main

func main() {
	nums := []int{2, 7, 11, 15}
	target := 7
	twoSum(nums, target)
}

//暴力解法
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target-nums[i] {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	var munIndexMap map[int]int
	//
	for i := range nums {
		val := nums[i]

		if index, ok := munIndexMap[target-val]; ok {
			return []int{i, index}
		}

		munIndexMap[val] = i
	}
	return []int{}
}

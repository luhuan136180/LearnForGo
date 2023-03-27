package _026__删除有序数组中的重复项

func removeDuplicates(nums []int) int {
	//使用双指针
	if len(nums) == 0 {
		return 0
	}
	j := 1

	//I，向前比对，快慢指针一开始在同一个位置
	for i := 1; i < len(nums)-1; i++ {
		//不一样时，插入
		if nums[i] != nums[i+1] {
			nums[j] = nums[i]
			j++
		}

	}

	return j
}

package _283__移动零

func moveZeroes(nums []int) {
	lenth := len(nums)

	slow := 0

	for fast := 0; fast < lenth; fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}

}

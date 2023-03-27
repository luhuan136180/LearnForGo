package _027__移除元素

//使用另一个数组
func removeElement(nums []int, val int) int {

	err := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[err] = nums[i]
			err++
		}
	}

	nums = nums[:err]
	return err
}

package main

//贪心的通解2代版本
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	curindex, nextDistance := 0, 0
	ans := 0
	for i := 0; i < len(nums)-1; i++ {
		nextDistance = max(nums[i]+i, nextDistance)
		if i == curindex {
			curindex = nextDistance
			ans++
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//贪心的通解一代版本
func jump2(nums []int) int {
	if len(nums) == 1 {
		return 0

	}
	//curDistance,nextDistance:标注的是下标，当前所能覆盖的最远位置的下标，和已知的下一步可达的最远覆盖位置的下标
	curDistance, nextDistance, ans := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		nextDistance = max(nums[i]+i, nextDistance)
		if i == curDistance {
			if curDistance != len(nums)-1 {
				ans++
				curDistance = nextDistance
				if nextDistance >= len(nums)-1 {
					break
				}
			} else {
				break
			}
		}
	}
	return ans
}

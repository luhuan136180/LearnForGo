package _454__四数相加_II

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	//key:a+b的数值，value:a+b数值出现的次数
	set := make(map[int]int, 0)
	// 遍历nums1和nums2数组，统计两个数组元素之和，和出现的次数，放到map中
	for _, a := range nums1 {
		for _, b := range nums2 {
			set[a+b]++
		}
	}
	count := 0
	// 遍历nums3和nums4数组，找到如果 0-(c+d) 在map中出现过的话，就把map中key对应的value也就是出现次数统计出来
	for _, c := range nums3 {
		for _, d := range nums4 {
			count += set[0-(c+d)]
		}
	}
	return count
}

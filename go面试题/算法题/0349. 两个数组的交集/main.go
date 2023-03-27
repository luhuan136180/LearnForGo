package _349__两个数组的交集

func intersection(nums1 []int, nums2 []int) []int {
	map1, map2 := make(map[int]int), make(map[int]int)
	ans := []int{}

	for _, val := range nums1 {
		map1[val]++
	}
	for _, val := range nums2 {
		map2[val]++

	}
	for key, _ := range map1 {
		if map2[key] != 0 {
			ans = append(ans, key)
		}
	}

	return ans
}
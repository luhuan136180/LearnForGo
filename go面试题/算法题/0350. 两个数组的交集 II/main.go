package _350__两个数组的交集_II

func intersect(nums1 []int, nums2 []int) []int {
	set := make(map[int]int, 0) //mapKey表示数组中出现的元素，val表示出现的geshu
	ans := make([]int, 0)

	//
	for _, val := range nums1 {
		set[val]++
	}

	for _, val := range nums2 {

		if _, ok := set[val]; ok {
			ans = append(ans, val)
			set[val]--

		}
		if set[val] == 0 {
			delete(set, val)
		}
	}

	return ans
}

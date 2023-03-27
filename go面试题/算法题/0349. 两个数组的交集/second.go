package _349__两个数组的交集

func intersection2(nums1 []int, nums2 []int) []int {
	//单个map的做法
	//一个map记录一个nums，然后遍历另一个nums，在map中查询，有结果则在结果集中添加
	set := make(map[int]struct{}, 0)
	ans := []int{}

	for _, val := range nums1 {
		if _, ok := set[val]; !ok {
			//map中没有录入该元素
			set[val] = struct{}{}
		}
	}

	for _, val := range nums2 {
		if _, ok := set[val]; ok {
			ans = append(ans, val)
			delete(set, val) //删除，防止重复
		}
	}

	return ans
}

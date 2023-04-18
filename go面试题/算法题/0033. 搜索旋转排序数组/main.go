package _033__搜索旋转排序数组

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	// 先找旋转点, 例如在nums=【4,5,6,7,1,2,3】中找到7，或在【1,2,3,4,5,6,7】中找到7
	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == nums[0] {
			l = mid + 1 // 继续向右找旋转点
		} else if nums[mid] > nums[0] {
			l = mid + 1 // 继续向右找旋转点
		} else if nums[mid] < nums[0] {
			r = mid - 1 // 继续向左找旋转点
		}
	}
	// 若有两段，如【4,5,6,7,1,2,3】：   则此时 l 是第二段首元素【1】, r 是第一段的尾元素【7】
	// 若只有一段，如【1,2,3,4,5,6,7】： 则此时 l是结尾之后【8】，r是结尾元素【7】

	// 判断target在旋转点的哪一侧
	if target >= nums[0] {
		l = 0 // 若target是在旋转点的左侧则在旋转点左侧二分搜索target，即在【4,5,6,7,1,2,3】中的【4,5,6,7】中搜索，或在【1,2,3,4,5,6,7】中的【1,2,3,4,5,6,7】中搜索
	} else {
		r = n - 1 // 在旋转点右侧二分搜索target，即在【1,2,3】中搜索
	}
	// 确定好有序区间了，直接二分搜索target即可
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1 // 继续向左找
		} else if nums[mid] < target {
			l = mid + 1 // 继续向右找
		}
	}
	return -1
}

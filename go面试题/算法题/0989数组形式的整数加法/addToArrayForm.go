package main

func addToArrayForm(num []int, k int) []int {
	var res []int
	carry := 0 //余数——进位数值
	l1 := len(num) - 1
	//遍历相加
	for l1 >= 0 || k != 0 {
		//初始化
		x, y := 0, 0
		if l1 >= 0 {
			x = num[l1]
		}
		if k != 0 {
			y = k % 10
		}

		sum := x + y + carry
		res = append(res, sum%10)
		carry = sum / 10

		l1--
		k /= 10
	}
	//加余数
	if carry != 0 {
		res = append(res, carry)
	}
	//反正切片
	left, right := 0, len(res)-1
	for left < right {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}
	return res
}

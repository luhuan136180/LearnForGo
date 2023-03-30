package main

func lemonadeChange3(bills []int) bool {
	//模拟手上有n，m张 5元，10元
	five, ten := 0, 0
	//遍历购买的数组
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			five++
		}
		if bills[i] == 10 {
			five--
			ten++
		}
		if bills[i] == 20 {
			if ten > 0 {
				ten--
				five -= 1
			} else {
				five -= 3
			}

		}
		if five < 0 && ten < 0 {
			return false
		}
	}
	return true
}

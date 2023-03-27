package _258__各位相加

func addDigits(num int) int {
	for {
		if num/10 == 0 {
			return num
		}
		num = getSum(num)
	}
}

func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10 //% 取余数
		n = n / 10
	}
	return sum
}

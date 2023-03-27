package _202__快乐数

func isHappy(n int) bool {
	m := make(map[int]bool)

	//n==1 或m【n】重复则退出
	for n != 1 && !m[n] {
		n, m[n] = getSum(n), true
	}
	return n == 1
}

func getSum(n int) int {
	sum := 0

	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}

package main

func main() {

}
func toSuml(num []int, target int) []int {
	len := len(num)
	for i := 0; i < len; i++ {
		val := num[i]
		for j := i + 1; j < len; j++ {
			if num[j] == target-val {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

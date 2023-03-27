package main

func numDifferentIntegers(word string) int {
	res := map[string]int{} //res用于存储word中存在的整数
	size := len(word)
	for i := 0; i < size; i++ {

		if word[i] >= '0' && word[i] <= '9' { //遇到数字
			for i < size && word[i] == '0' {
				i++
			} //找到不为0的其实数字
			j := i
			for j < size && word[j] >= '0' && word[j] <= '9' {
				j++
			}
			res[word[i:j]] = 1
			i = j
		}
	}
	return len(res)
}

package _242__有效的字母异位词

func isAnagram3(s string, t string) bool {
	arr := [26]int{}

	for _, val := range s {
		arr[val-rune('a')]++
	}

	for _, val := range t {
		arr[val-rune('a')]--
	}

	return arr == [26]int{}
}

package _242__有效的字母异位词

func isAnagram(s string, t string) bool {
	mapS, mapT := make(map[byte]int), make(map[byte]int)

	for i := 0; i < len(s); i++ {
		mapS[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		mapT[t[i]]++
	}

	for key, val := range mapS {
		if mapT[key] != val {
			return false
		}
	}
	for key, val := range mapT {
		if mapS[key] != val {
			return false
		}
	}
	return true
}
func isAnagram2(s string, t string) bool {
	record := [26]int{}

	for _, r := range s {
		record[r-rune('a')]++
	}
	for _, r := range t {
		record[r-rune('a')]--
	}
	// record数组如果有的元素不为零0，说明字符串s和t 一定是谁多了字符或者谁少了字符。
	return record == [26]int{}
}

package _383__赎金信

//map做法
func canConstruct(ransomNote string, magazine string) bool {
	//K:元素，val:该字符出现的个数
	set := make(map[byte]int, 0)

	for i := 0; i < len(magazine); i++ {
		set[magazine[i]]++
	}

	for j := 0; j < len(ransomNote); j++ {
		if _, ok := set[ransomNote[j]]; ok {
			set[ransomNote[j]]--
			if set[ransomNote[j]] == 0 {
				delete(set, ransomNote[j])
			}
		} else {
			return false
		}
	}

	return true
}

//数组做法
func canConstruct2(ransomNote string, magazine string) bool {
	record := make([]int, 26)

	for _, v := range magazine {
		record[v-'a']++
	}

	for _, v := range ransomNote {
		record[v-'a']--
		if record[v-'a'] < 0 {
			return false
		}
	}
	return true
}

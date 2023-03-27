package _049__字母异位词分组

func groupAnagrams(strs []string) [][]string {
	Smap := map[[26]int][]string{}

	for _, val := range strs {
		temp := [26]int{}
		for _, b := range val {
			temp[b-rune('a')]++
		}
		Smap[temp] = append(Smap[temp], val)
	}

	ans := make([][]string, 0, len(Smap))
	for _, v := range Smap {
		ans = append(ans, v)
	}

	return ans

}

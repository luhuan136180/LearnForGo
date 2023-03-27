package main

import "fmt"

func test_3_12(n1, n2 string) bool {
	m := make(map[rune]int)
	for _, val := range n1 {
		m[val]++
	}
	for _, val := range n2 {
		m[val]--
	}
	for _, val := range m {
		if val != 0 {
			return false
		}
	}
	return true
}

func main() {
	s1, s2 := "sdfghjkuytfvb", "kjhgfdstyubvf"
	fmt.Printf("first string is :%s\nsecond string is :%s\n", s1, s2)
	if test_3_12(s1, s2) {
		fmt.Println("them have same characters")
	} else {
		fmt.Printf("they don't have different characters")
	}
}

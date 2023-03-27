package main

import "strings"

func basename2(s string) string {
	//子串sep在字符串s中最后一次出现的位置，不存在则返回-html。
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

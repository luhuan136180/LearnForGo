package main

import "strconv"

func Sprint(x interface{}) string {
	type stringer interface {
		string() string //func (stringer) string() string
	}
	switch x := x.(type) {
	case stringer:
		return x.string()
	case string:
		return x
	case int:
		return strconv.Itoa(x) //func Itoa(i int) string
	case bool:
		if x {
			return "true"
		}
		return "false"
	default:
		return "???"
	}
}

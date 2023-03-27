package main

import (
	"fmt"
	"sort"
)

//使用map来记录提交相同的字符串列表的次数
var m = make(map[string]int)

func k(list []string) string { return fmt.Sprintf("%q", list) }

func Add(list []string)       { m[k(list)]++ }
func Count(list []string) int { return m[k(list)] }

func main() {
	ages := make(map[string]int)

	ages["alice"] = 31
	ages["charlie"] = 34
	fmt.Println(ages["alice"])
	//删除元素
	delete(ages, "alice")
	//	map中的元素并不是一个变量，因此我们不能对map的元素进行取址操作
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	//	sort
	//var names []string
	names := make([]string, 0, len(ages))

	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	//

}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

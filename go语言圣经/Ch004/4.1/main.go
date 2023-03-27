package main

import "fmt"

//数组的每个元素都被初始化为元素类型对应的零值
func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, v := range a {
		fmt.Printf("%d  %d \n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
	//var q [3]int = [3]int{html, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2]) // "0"

	//在数组的长度位置出现的是“...”省略号，
	//则表示数组的长度是根据初始化值的个数来计算
	q := [...]int{1, 2, 3} //根据初始值个数计算长度
	fmt.Printf("%T\n", q)  // "[3]int"
	m := [...]int{99: -1}
	fmt.Println(m)
	fmt.Println(len(m))

	a1 := [2]int{1, 2}
	b1 := [...]int{1, 2}
	c1 := [2]int{1, 3}
	fmt.Println(a1 == b1, a1 == c1, b1 == c1)
	// "true false false"
	d := [3]int{1, 2}
	fmt.Println(a == d)
	// compile error: cannot compare [2]int == [3]int

}

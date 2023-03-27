package main

import "fmt"

func main() {

	arr := &[6]int{3, 9, 2, 4, 7, 1}
	SelectSort(arr)

}
func SelectSort(arr *[6]int) {
	for i := 0; i < len(arr); i++ {
		min := arr[i]
		minindex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < min {
				min = arr[j]
				minindex = j
			}
		}
		if minindex != i {
			temp := arr[i]
			arr[i] = min
			arr[minindex] = temp
		}
		fmt.Printf("第%d次改动，%v\n", i+1, *arr)
	}

}

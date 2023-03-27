package main

import "fmt"

func main() {

	arr := &[6]int{3, 9, 2, 4, 7, 1}
	SelectSort(arr)

}
func SelectSort(arr *[6]int) {
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		tempIndex := i - 1
		for tempIndex >= 0 && arr[tempIndex] < temp {
			arr[tempIndex+1] = arr[tempIndex]
			tempIndex--
		}
		if tempIndex+1 != i {
			arr[tempIndex+1] = temp
		}
		fmt.Printf("第%d次插入后%v\n", i, *arr)
	}
}

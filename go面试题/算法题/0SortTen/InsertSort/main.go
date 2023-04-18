package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var arr [8]int
	for i := 0; i < 8; i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)
	InsertSort(&arr)
	fmt.Println(arr)
}

func InsertSort(arr *[8]int) {
	for i := 1; i < len(arr); i++ {
		indexval := arr[i]
		index := i - 1
		for index >= 0 && arr[index] >= indexval {
			arr[index+1] = arr[index]
			index--
		}
		if index+1 != i {
			arr[index+1] = indexval
		}
	}
}

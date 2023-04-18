package main

func bubbleSort(arr []int, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = tmp
			}
		}
	}
}

//优化1.0
func bubbleSort2(arr []int, n int) {
	flag := false
	for i := 0; i < n; i++ {
		flag = false
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				flag = true
				tmp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = tmp
			}
		}
		if !flag {
			//说明没有交换，则表明【0，len-1-i】已经是有序的了
			break
		}
	}
}

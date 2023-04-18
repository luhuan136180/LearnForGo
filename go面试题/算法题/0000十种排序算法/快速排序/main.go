package 快速排序

import "fmt"

/*
1、选取第一个数为基准

2、将比基准小的数交换到前面，比基准大的数交换到后面

3、对左右区间重复第二步，直到各区间只有一个数

我们从数组中选择一个元素，我们把这个元素称之为中轴元素吧，然后把数组中所有小于中轴元素的元素放在其左边，所有大于或等于中轴元素的元素放在其右边，显然，此时中轴元素所处的位置的是有序的。也就是说，我们无需再移动中轴元素的位置。

从中轴元素那里开始把大的数组切割成两个小的数组(两个数组都不包含中轴元素)，接着我们通过递归的方式，让中轴元素左边的数组和右边的数组也重复同样的操作，直到数组的大小为1，此时每个元素都处于有序的位置。

*/
func QuickSort(arr []int) []int {
	if len(arr) < 2 { // 如果数组长度小于2，那么该数组已经有序，直接返回
		return arr
	} else {
		pivotIndex := 0                 // 设置pivotIndex，取数组第一个元素作为基准（pivot）
		for i := 1; i < len(arr); i++ { // 从数组第二个元素开始遍历
			if arr[i] < arr[pivotIndex] { // 如果该元素小于基准，就将该元素与pivotIndex+1位置处的元素交换，并将pivotIndex+1
				arr[i], arr[pivotIndex+1] = arr[pivotIndex+1], arr[i]
				pivotIndex++
			}
		}
		arr[0], arr[pivotIndex] = arr[pivotIndex], arr[0]      // 基准归位
		left := QuickSort(arr[:pivotIndex])                    //对左侧子数组递归调用快速排序
		right := QuickSort(arr[pivotIndex+1:])                 //对右侧子数组递归调用快速排序
		return append(append(left, arr[pivotIndex]), right...) // 返回排序后的数组
	}
}

func main() {
	arr := []int{1, 11, 4, 8, 0, 7, 6, 2, 3, 10, 5, 9} // 待排序数组
	arr = QuickSort(arr)                               // 快速排序
	fmt.Println(arr)                                   // 输出排序后的数组
}

func QuickSort2(arr []int) []int {
	if len(arr) < 2 {
		//只有一个元素或者没有元素
		return arr
	} else {
		//有多个元素
		index := 0                      //设置基准元素下标,取本数组第一个元素为基准元素
		for i := 1; i < len(arr); i++ { //从本数组第二个元素开始遍历
			if arr[i] < arr[index] { //如果该元素小于基准，则将该元素与index+1位置的元素互换位置，并将基准元素洗标+1
				arr[i], arr[index+1] = arr[index+1], arr[i] //交换
				index++                                     //基准下标后移，每有一个小于它的元素，就将下标后移，下标以前的就是下雨基准的元素
			}
		}
		//遍历完成，所有小于基准元素的元素，已经移动至index下标之前，现在要将基准元素移动到基准下标的位置
		arr[0], arr[index] = arr[index], arr[0] //基准元素归为,此时基准元素是有序的
		//对前后分别进行快速排序
		left := QuickSort(arr[:index])
		right := QuickSort(arr[index+1:])
		return append(append(left, arr[index]), right...)
	}
}
func QuickSort3(arr []int) []int {
	if len(arr) < 2 {
		return arr //有序的
	} else {
		index := 0
		key := arr[index]
		for i := 1; i < len(arr); i++ {
			if arr[i] < key {
				arr[i], arr[index+1] = arr[index+1], arr[i]
				index++
			}
		}
		arr[0], arr[index] = arr[index], arr[0]
		left := QuickSort(arr[:index])
		right := QuickSort(arr[index+1:])
		return append(append(left, arr[index]), right...)
	}

}

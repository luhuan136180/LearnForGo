package main

type KthLargest2 struct {
	size  int //堆的大小，不包括数组第一个元素
	data  []int
	count int //当前元素数量
}

func Countructor(k int, nums []int) KthLargest2 {
	kth := KthLargest2{}
	kth.size = k
	kth.data = []int{0}
	for _, num := range nums {
		kth.Add(num)
	}
	return kth
}

func (kl *KthLargest2) Add(val int) int {
	if kl.count < kl.size-1 { //模拟下标
		kl.data = append(kl.data, val) //直接添加进末尾
		kl.count += 1
	} else if kl.count == kl.size-1 { //还差一个填满
		kl.data = append(kl.data, val)
		kl.count += 1 //填满
		//第一次填满,是堆有序
		n := len(kl.data) - 1 //data的最大下标
		for i := n / 2; i > 1; i-- {
			heapify(kl.data, i)
		}

	} else {

	}
	return kl.data[1]
}

//heapify 从给定的i向下堆化为小顶堆
func heapify(a []int, i int) {
	n := len(a) - 1
	for {
		minPos := i
		if i*2 <= n && a[i*2] < a[minPos] {
			minPos = i * 2
		}
		if i*2+1 <= n && a[i*2+1] < a[minPos] {
			minPos = i*2 + 1
		}
		if minPos == i {
			break
		}
		a[minPos], a[i] = a[i], a[minPos]
		i = minPos
	}
}

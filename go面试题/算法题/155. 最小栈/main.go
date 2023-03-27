package _55__最小栈

import "math"

//设计一个辅助栈，每当入栈元素时，计算包括入栈元素在内的最小值，将其压入辅助栈
type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(val int) {
	//得到辅助栈中当前最小元素的大小
	//比较，并将新的最小值压入辅助栈
	//将数据压入栈中
	top := this.minStack[len(this.minStack)-1]
	top = min(val, top)
	this.stack = append(this.stack, val)
	this.minStack = append(this.minStack, top)
}

func (this *MinStack) Pop() {
	//val:=this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]

}

func (this *MinStack) Top() int {
	val := this.stack[len(this.stack)-1]
	return val
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

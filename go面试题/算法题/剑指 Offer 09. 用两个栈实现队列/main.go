package 剑指_Offer_09__用两个栈实现队列

//入队栈和出队栈
type CQueue struct {
	InStack  []int
	OutStack []int
}

func Constructor() CQueue {
	return CQueue{
		InStack:  []int{},
		OutStack: []int{},
	}
}

func (this *CQueue) AppendTail(value int) {
	this.InStack = append(this.InStack, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.OutStack) == 0 && len(this.InStack) == 0 {
		return -1
	}
	if len(this.OutStack) != 0 {
		val := this.OutStack[len(this.OutStack)-1]
		this.OutStack = this.OutStack[:len(this.OutStack)-1]
		return val
	} else {
		for len(this.InStack) != 0 {
			top := this.InStack[len(this.InStack)-1]
			this.InStack = this.InStack[:len(this.InStack)-1]
			this.OutStack = append(this.OutStack, top)
		}
		val := this.OutStack[len(this.OutStack)-1]
		this.OutStack = this.OutStack[:len(this.OutStack)-1]
		return val
	}

}

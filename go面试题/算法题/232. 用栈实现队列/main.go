package main

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

type MyQueue struct {
	in, out []int //in:输入宅    out:输出栈
}

func Constructor() MyQueue {
	return MyQueue{
		in:  make([]int, 0),
		out: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

func (this *MyQueue) Pop() int {
	//当out栈中有值时，不能讲in栈中的值放入out中只有out中为空，才将in放入
	if len(this.out) == 0 {
		for len(this.in) != 0 {
			val := this.in[len(this.in)-1]
			this.out = append(this.out, val)
			this.in = this.in[:len(this.in)-1]
		}
	}
	result := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return result
}

func (this *MyQueue) Peek() int {
	if len(this.out) == 0 {
		for len(this.in) != 0 {
			val := this.in[len(this.in)-1]
			this.out = append(this.out, val)
			this.in = this.in[:len(this.in)-1]
		}
	}
	return this.out[len(this.out)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.in) == 0 && len(this.out) == 0
}

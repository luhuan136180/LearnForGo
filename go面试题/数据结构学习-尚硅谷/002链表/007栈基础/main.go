package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	MaxTop int    //表示栈的容量
	Top    int    //表示栈顶
	arr    [5]int //数组模拟
}

func (this *Stack) Push(val int) (err error) {
	//判断是否栈满了
	if this.Top == this.MaxTop-1 {
		fmt.Println("Stack full")
		return errors.New("stack full")
	}

	this.Top++
	this.arr[this.Top] = val
	return
}

//出栈
func (this *Stack) Pop() (val int, err error) {
	if this.Top == -1 {
		fmt.Println("stach empty")
		return 0, errors.New("stack empty")
	}
	val = this.arr[this.Top]
	this.Top--
	return val, nil
}

func (this *Stack) List() {
	if this.Top == -1 {
		fmt.Println("Stack empty")
		return
	}
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, this.Top)
	}
}

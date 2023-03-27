package main

import (
	"fmt"
	"sync"
)

var pool *sync.Pool

type Person struct {
	Name string
}

func initPool() {
	var a = func() interface{} {
		fmt.Println("Createing a new Person")
		return new(Person)
	}
	pool = &sync.Pool{
		New: a,
	}
}

func main() {
	initPool()
	p := pool.Get().(*Person)
	fmt.Println("首次从pool里获取：", p)

	p.Name = "first"
	fmt.Printf("设置p.Name = %s\n", p.Name)

	pool.Put(p)

	fmt.Println("Pool 里已经有一个对象：&{first},调用Get：", pool.Get().(*Person))
	fmt.Println("pool里没有对象，调用get", pool.Get().(*Person))
}

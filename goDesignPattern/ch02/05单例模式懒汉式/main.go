package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type singelton struct{}

//定义一个互斥锁
var lock sync.Mutex

//标记
var initialized uint32

var instance *singelton

func GetInstance() *singelton {

	if atomic.LoadUint32(&initialized) == 1 {
		return instance //已经存在
	}

	//
	lock.Lock()
	defer lock.Unlock()

	//只有首次GetInstance()方法被调用，才会生成这个单例的实例
	//当并发时，可能会重复创建
	if instance == nil {
		instance = new(singelton)
		atomic.StoreUint32(&initialized, 1)
	}

	//接下来的GetInstance直接返回已经申请的实例即可
	return instance
}

func (s *singelton) SomeThing() {
	fmt.Println("单例对象的某方法")
}

func main() {
	s := GetInstance()
	s.SomeThing()
}

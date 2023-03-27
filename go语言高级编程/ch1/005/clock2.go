package main

import "sync"

type singleton struct{}

//
//var (
//	instance    *singleton
//	initialized uint32
//	mu          sync.Mutex
//)
//
//func Instance() *singleton {
//	if atomic.LoadUint32(&initialized) == 1 {
//		return instance
//	}
//
//	mu.Lock()
//	defer mu.Unlock()
//
//	if instance == nil {
//		defer atomic.StoreUint32(&initialized, 1)
//		instance = &singleton{}
//	}
//	return instance
//}

//一下为sync。One标准库的实现中
//type Once struct {
//	m    Mutex
//	done uint32
//}
//
//func (o *Once) Do(f func()) {
//	if atomic.LoadUint32(&o.done) == 1 {
//		return
//	}
//
//	o.m.Lock()
//	defer o.m.Unlock()
//
//	if o.done == 0 {
//		defer atomic.StoreUint32(&o.done, 1)
//		f()
//	}
//}

var (
	instance *singleton
	once     sync.Once
)

func Instance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

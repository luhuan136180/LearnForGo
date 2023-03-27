package __3

import "sync"

//允许多个只读操作并行执行，但写操作会完全互斥。
var mu sync.RWMutex
var balance int

func Balance() int {
	mu.RLock() // readers lock
	defer mu.RUnlock()
	return balance
}
func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()          //使用defer，固定函数返回时调用
	balance = balance + amount //临界区

}

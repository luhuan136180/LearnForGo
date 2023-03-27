package bank3

import "sync"

var (
	balance int
	mu      sync.Mutex
)

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()          //使用defer，固定函数返回时调用
	balance = balance + amount //临界区

}

func Balance() int {
	mu.Lock()
	//b := balance
	////mu.Unlock()
	//return b

	defer mu.Unlock()
	return balance //Unlock会在return语句读取完
	// balance的值之后执行
	//	不再需要一个本地变量b了
}

//可能导致并发支付时，出现问题；例：同时买车跟买水，导致都无法执行
//解决思路见blank4

func Withdraw(amount int) bool {
	Deposit(-amount)
	if Balance() < 0 {
		Deposit(amount)
		return false
	}
	return true
}

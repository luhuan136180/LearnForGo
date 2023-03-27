package main

//基于select与 time.After的超市判断代码
//select {
//case v := <-in:
//fmt.Println(v)
//case <-time.After(time.Second):
//return // 超时
//}

//select 的 default 分支实现非阻塞的管道发送或接收操作
//select {
//case v := <-in:
//fmt.Println(v)
//default:
//// 没有数据
//}

//阻止main的退出
//func main() {
//	// do some thins
//	select{}
//}

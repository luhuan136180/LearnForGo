package main

func main() {

	//func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
	//	c := newCancelCtx(parent)
	//	propagateCancel(parent, &c)
	//	return &c, func() { c.cancel(true, Canceled) }
	//}

	//context.newCancelCtx 将传入的上下文包装成私有结构体 context.cancelCtx；
	//context.propagateCancel 会构建父子上下文之间的关联，当父上下文被取消时，子上下文也会被取消：

}

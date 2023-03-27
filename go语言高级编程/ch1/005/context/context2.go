package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//果当前函数没有上下文作为入参，
	//我们都会使用 context.Background 作为起始的上下文向下传递。
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with ", duration)
	}
}

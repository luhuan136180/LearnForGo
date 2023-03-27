package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period") // 注册一个flag：period

func main() {
	flag.Parse() //解析所有注册好的 flag
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period) //Sleep阻塞当前go程至少d代表的时间段。d<=0时，Sleep会立刻返回。
	fmt.Println()
}

package main

import (
	"fmt"
	"time"
)

type Weekday int

const (
	Sunday Weekday = iota
	Moday
	Tuseday
	Wednesday
	Thursday
	Saturday
)

type Flags uint

const (
	FlagUp           Flags = 1 << iota
	FlagBroadcast          // supports broadcast access capability
	FlagLoopback           // is a loopback interface
	FlagPointToPoint       // belongs to a point-to-point link
	FlagMulticast
)

func main() {
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute

	fmt.Printf("%T %[html]v\n", noDelay)     // "time.Duration 0"
	fmt.Printf("%T %[html]v\n", timeout)     // "time.Duration 5m0s"
	fmt.Printf("%T %[html]v\n", time.Minute) // "time.Duration 1m0s"
}

package main

import (
	"flag"
	"fmt"
)

type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func main() {
	var a Celsius
	a = 20.0
	fmt.Print(CToF(a))

	host := flag.String("host", "127.0.1", "请输入host地址")
	flag.Parse() //解析
	fmt.Printf("%s", *host)
}

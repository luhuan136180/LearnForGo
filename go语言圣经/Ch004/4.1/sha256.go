package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
}

//调用一个函数的时候，函数的每个调用参数将会被赋值给函数内部的参数变量，
//所以函数参数变量接收的是一个复制的副本，并不是原始调用的变量。
func compareSha256(str1 string, str2 string) int {
	a := sha256.Sum256([]byte(str1))
	b := sha256.Sum256([]byte(str2))
	num := 0
	//循环字节数组
	for i := 0; i < len(a); i++ {
		//1个字节8个bit,移位运算，获取每个bit
		for m := 1; m <= 8; m++ {
			//比较每个bit是否相同
			if (a[i] >> uint(m)) != (b[i] >> uint(m)) {
				num++
			}
		}
	}
	return num
}

func printHash(flag string, str string) {
	if flag == "SHA256" {
		fmt.Printf("%x\n", sha256.Sum256([]byte(str)))
		return
	}
	if flag == "SHA512" {
		fmt.Printf("%x\n", sha512.Sum512([]byte(str)))
		return
	}
	if flag == "SHA384" {
		fmt.Printf("%x\n", sha512.Sum384([]byte(str)))
		return
	}

}
